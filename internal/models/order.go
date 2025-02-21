package models

type Order struct {
	ID                 uint    `json:"id" gorm:"primaryKey"`
	StoreID            int     `json:"store_id" binding:"required"`
	MerchantOrderID    string  `json:"merchant_order_id"`
	RecipientName      string  `json:"recipient_name" binding:"required"`
	RecipientPhone     string  `json:"recipient_phone" binding:"required"`
	RecipientAddress   string  `json:"recipient_address" binding:"required"`
	RecipientCity      int     `json:"recipient_city"`
	RecipientZone      int     `json:"recipient_zone"`
	RecipientArea      int     `json:"recipient_area"`
	DeliveryType       int     `json:"delivery_type"`
	ItemType           int     `json:"item_type"`
	SpecialInstruction string  `json:"special_instruction"`
	ItemQuantity       int     `json:"item_quantity" binding:"required"`
	ItemWeight         float64 `json:"item_weight" binding:"required"`
	AmountToCollect    float64 `json:"amount_to_collect" binding:"required"`
	ItemDescription    string  `json:"item_description"`
	DeliveryFee        float64 `json:"delivery_fee"`
	CODFee             float64 `json:"cod_fee"`
	OrderStatus        string  `json:"order_status" gorm:"default:'Pending'"`
}
