package db

import (
	"fmt"
	"github.com/filipe-rds/microservices/shipping/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ShippingItem struct {
	gorm.Model
	ShippingId  uint   `gorm:"not null"`
	ProductCode string `gorm:"not null"`
	Quantity    int32  `gorm:"not null"`
}

type Shipping struct {
	gorm.Model
	OrderId      int64          `gorm:"not null"`
	DeliveryDays int32          `gorm:"not null"`
	CreatedAt    int64          `gorm:"not null"`
	Items        []ShippingItem `gorm:"foreignKey:ShippingId"`
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %v", err)
	}
	
	err = db.AutoMigrate(&Shipping{}, &ShippingItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	
	return &Adapter{db: db}, nil
}

func (a *Adapter) Get(id string) (domain.Shipping, error) {
	var shippingEntity Shipping
	err := a.db.Preload("Items").First(&shippingEntity, id).Error
	if err != nil {
		return domain.Shipping{}, err
	}
	
	return toShippingDomain(shippingEntity), nil
}

func (a *Adapter) Save(shipping *domain.Shipping) error {
	shippingEntity := toShippingEntity(*shipping)
	
	result := a.db.Create(&shippingEntity)
	if result.Error != nil {
		return result.Error
	}
	
	shipping.ID = int64(shippingEntity.Model.ID)
	return nil
}

func toShippingDomain(shipping Shipping) domain.Shipping {
	var items []domain.ShippingItem
	for _, item := range shipping.Items {
		items = append(items, domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}
	
	return domain.Shipping{
		ID:           int64(shipping.Model.ID),
		OrderId:      shipping.OrderId,
		Items:        items,
		DeliveryDays: shipping.DeliveryDays,
		CreatedAt:    shipping.CreatedAt,
	}
}

func toShippingEntity(shipping domain.Shipping) Shipping {
	var items []ShippingItem
	for _, item := range shipping.Items {
		items = append(items, ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}
	
	return Shipping{
		OrderId:      shipping.OrderId,
		DeliveryDays: shipping.DeliveryDays,
		CreatedAt:    shipping.CreatedAt,
		Items:        items,
	}
}
