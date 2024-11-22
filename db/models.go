package db

import "time"

type Invoice struct {
	InvoiceNumber   int64     `json:"invoice_number"`
	CustomerName    string    `json:"customer_name"`
	CustomerEmail   string    `json:"customer_email"`
	CustomerPhone   string    `json:"customer_phone"`
	CustomerAddress string    `json:"customer_address"`
	SenderName      string    `json:"sender_name"`
	SenderEmail     string    `json:"sender_email"`
	SenderPhone     string    `json:"sender_phone"`
	SenderAddress   string    `json:"sender_address"`
	IssueDate       string    `json:"issue_date"`
	DueDate         string    `json:"due_date"`
	Status          string    `json:"status"`
	Subtotal        string    `json:"subtotal"`
	DiscountRate    string    `json:"discount_rate"`
	Discount        string    `json:"discount"`
	TotalAmount     string    `json:"total_amount"`
	PaymentInfo     string    `json:"payment_info"`
	BillingCurrency string    `json:"billing_currency"`
	Note            string    `json:"note"`
	CreatedAt       time.Time `json:"created_at"`
}

type LineItem struct {
	ID            int64  `json:"id"`
	InvoiceNumber int64  `json:"invoice_number"`
	Description   string `json:"description"`
	Quantity      int64  `json:"quantity"`
	UnitPrice     string `json:"unit_price"`
	TotalPrice    string `json:"total_price"`
}
