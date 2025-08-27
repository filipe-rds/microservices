package domain

import "time"

type ShippingItem struct {
	ProductCode string `json:"product_code"`
	Quantity    int32  `json:"quantity"`
}

type Shipping struct {
	ID           int64          `json:"id"`
	OrderId      int64          `json:"order_id"`
	Items        []ShippingItem `json:"items"`
	DeliveryDays int32          `json:"delivery_days"`
	CreatedAt    int64          `json:"created_at"`
}

func NewShipping(orderId int64, items []ShippingItem) Shipping {
	return Shipping{
		OrderId:      orderId,
		Items:        items,
		DeliveryDays: calculateDeliveryDays(items),
		CreatedAt:    time.Now().Unix(),
	}
}

// calculateDeliveryDays calcula o prazo de entrega baseado na quantidade total de itens
// Prazo mínimo é de 1 dia, e a cada 5 unidades é adicionado um dia a mais
func calculateDeliveryDays(items []ShippingItem) int32 {
	totalQuantity := int32(0)
	for _, item := range items {
		totalQuantity += item.Quantity
	}
	
	// Prazo mínimo de 1 dia + 1 dia a cada 5 unidades
	deliveryDays := int32(1) + (totalQuantity / 5)
	
	return deliveryDays
}
