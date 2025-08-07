package domain

import "time"

type OrderItem struct {
	ProductCode string `json:"product_code"`
	UnitPrice float32 `json:"unit_price"`
	Quantity int32 `json:"quantity"`
}

type Order struct {
	ID int64 `json:"id"`
	CostumerID int64 `json:"costumer_id"`
	Status string `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt int64 `json:"created_at"`
}

func NewOrder(costumerId int64, orderItems []OrderItem) Order{
	return Order{
		CreatedAt: time.Now().Unix(),
		Status: "Pending",
		CostumerID: costumerId,
		OrderItems: orderItems,
	}
}

func (o *Order) TotalPrice() float32 {
	var totalPrice float32
	for _, orderItem := range o.OrderItems{
		totalPrice += orderItem.UnitPrice * float32(orderItem.Quantity)
	}
	return totalPrice
}

