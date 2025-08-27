package db

import (
	"fmt"

	"github.com/filipe-rds/microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerId int64
	Status string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice float32
	Quantity int32
	OrderId uint
}

type Product struct {
	gorm.Model
	ProductCode   string  `gorm:"uniqueIndex;not null"`
	Name          string  `gorm:"not null"`
	Price         float32 `gorm:"not null"`
	StockQuantity int     `gorm:"not null;default:0"`
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, OpenErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if OpenErr != nil {
		return nil, fmt.Errorf("db connetction error: %v", OpenErr)
	}
	return &Adapter{db: db}, nil
}

func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)
	var orderItems []domain.OrderItem

	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice: orderItem.UnitPrice,
			Quantity: orderItem.Quantity,
		})
	}
	order := domain.Order{
		ID: int64(orderEntity.ID),
		CustomerId: orderEntity.CustomerId,
		OrderItems: orderItems,
		CreatedAt: orderEntity.CreatedAt.UnixNano(),
	}
	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) error{
	var orderItems []OrderItem
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice: orderItem.UnitPrice,
			Quantity: orderItem.Quantity,
		})
	}
	orderModel := Order{
		CustomerId: order.CustomerId,
		Status: order.Status,
		OrderItems: orderItems,
	}
	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}
	return res.Error
}

func (a Adapter) ValidateProducts(productCodes []string) error {
	var count int64
	err := a.db.Model(&Product{}).Where("product_code IN ?", productCodes).Count(&count).Error
	if err != nil {
		return fmt.Errorf("error validating products: %v", err)
	}
	
	if int(count) != len(productCodes) {
		return fmt.Errorf("one or more products do not exist in the database")
	}
	
	return nil
}


