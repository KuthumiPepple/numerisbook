package db

import (
	"context"
	"testing"
	"time"

	"github.com/kuthumipepple/numerisbook/util"
	"github.com/stretchr/testify/require"
)

func insertRandomInvoice(t *testing.T) Invoice {
	arg := InsertInvoiceParams{
		CustomerName:    util.RandomName(),
		CustomerEmail:   util.RandomEmail(),
		CustomerPhone:   util.RandomPhone(),
		CustomerAddress: util.RandomAddress(),
		SenderName:      util.RandomName(),
		SenderEmail:     util.RandomEmail(),
		SenderPhone:     util.RandomPhone(),
		SenderAddress:   util.RandomAddress(),
		IssueDate:       time.Now().Format(time.DateOnly),
		DueDate:         time.Now().AddDate(0, 0, 30).Format(time.DateOnly),
		Status:          util.RandomStatus(),
		Subtotal:        util.RandomMoney(),
		DiscountRate:    util.RandomRate(),
		Discount:        util.RandomMoney(),
		TotalAmount:     util.RandomMoney(),
		PaymentInfo:     util.RandomString(10),
		Note:            util.RandomString(20),
	}

	invoice, err := testStore.InsertInvoice(context.Background(), arg)
	require.NoError(t, err)

	require.NotZero(t, invoice.InvoiceNumber)

	require.Equal(t, arg.CustomerName, invoice.CustomerName)
	require.Equal(t, arg.CustomerEmail, invoice.CustomerEmail)
	require.Equal(t, arg.CustomerPhone, invoice.CustomerPhone)
	require.Equal(t, arg.CustomerAddress, invoice.CustomerAddress)
	require.Equal(t, arg.SenderName, invoice.SenderName)
	require.Equal(t, arg.SenderEmail, invoice.SenderEmail)
	require.Equal(t, arg.SenderPhone, invoice.SenderPhone)
	require.Equal(t, arg.SenderAddress, invoice.SenderAddress)
	require.Equal(t, arg.IssueDate, invoice.IssueDate)
	require.Equal(t, arg.DueDate, invoice.DueDate)
	require.Equal(t, arg.Status, invoice.Status)
	require.Equal(t, arg.Subtotal, invoice.Subtotal)
	require.Equal(t, arg.DiscountRate, invoice.DiscountRate)
	require.Equal(t, arg.Discount, invoice.Discount)
	require.Equal(t, arg.TotalAmount, invoice.TotalAmount)
	require.Equal(t, arg.PaymentInfo, invoice.PaymentInfo)
	require.Equal(t, "USD", invoice.BillingCurrency)

	require.Equal(t, arg.Note, invoice.Note)
	require.NotZero(t, invoice.CreatedAt)

	return invoice
}

func TestInsertInvoice(t *testing.T) {
	insertRandomInvoice(t)
}

func TestInsertLineItem(t *testing.T) {
	invoice := insertRandomInvoice(t)

	arg := InsertLineItemParams{
		InvoiceNumber: invoice.InvoiceNumber,
		Description:   util.RandomString(10),
		Quantity:      util.RandomInt(1, 10),
		UnitPrice:     util.RandomMoney(),
		TotalPrice:    util.RandomMoney(),
	}

	lineItem, err := testStore.InsertLineItem(context.Background(), arg)
	require.NoError(t, err)

	require.NotZero(t, lineItem.ID)
	require.Equal(t, arg.InvoiceNumber, lineItem.InvoiceNumber)
	require.Equal(t, arg.Description, lineItem.Description)
	require.Equal(t, arg.Quantity, lineItem.Quantity)
	require.Equal(t, arg.UnitPrice, lineItem.UnitPrice)
	require.Equal(t, arg.TotalPrice, lineItem.TotalPrice)
}

func TestCreateInvoiceTx(t *testing.T) {
	n := 10
	items := make([]LineItemsParam, n)
	for i := 0; i < n; i++ {
		items[i] = LineItemsParam{
			Description: util.RandomString(10),
			Quantity:    util.RandomInt(1, 10),
			UnitPrice:   util.RandomMoney(),
			TotalPrice:  util.RandomMoney(),
		}
	}

	arg := CreateInvoiceTxParams{
		CustomerName:    util.RandomName(),
		CustomerEmail:   util.RandomEmail(),
		CustomerPhone:   util.RandomPhone(),
		CustomerAddress: util.RandomAddress(),
		SenderName:      util.RandomName(),
		SenderEmail:     util.RandomEmail(),
		SenderPhone:     util.RandomPhone(),
		SenderAddress:   util.RandomAddress(),
		IssueDate:       time.Now().Format(time.DateOnly),
		DueDate:         time.Now().AddDate(0, 0, 30).Format(time.DateOnly),
		Status:          util.RandomStatus(),
		Subtotal:        util.RandomMoney(),
		DiscountRate:    util.RandomRate(),
		Discount:        util.RandomMoney(),
		TotalAmount:     util.RandomMoney(),
		PaymentInfo:     util.RandomString(10),
		Note:            util.RandomString(20),
		Items:           items,
	}

	result, err := testStore.CreateInvoiceTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, result.InvoiceNumber)
	require.Equal(t, arg.CustomerName, result.CustomerName)
	require.Equal(t, arg.CustomerEmail, result.CustomerEmail)
	require.Equal(t, arg.CustomerPhone, result.CustomerPhone)
	require.Equal(t, arg.CustomerAddress, result.CustomerAddress)
	require.Equal(t, arg.SenderName, result.SenderName)
	require.Equal(t, arg.SenderEmail, result.SenderEmail)
	require.Equal(t, arg.SenderPhone, result.SenderPhone)
	require.Equal(t, arg.SenderAddress, result.SenderAddress)
	require.Equal(t, arg.IssueDate, result.IssueDate)
	require.Equal(t, arg.DueDate, result.DueDate)
	require.Equal(t, arg.Status, result.Status)
	require.Equal(t, arg.Subtotal, result.Subtotal)
	require.Equal(t, arg.DiscountRate, result.DiscountRate)
	require.Equal(t, arg.Discount, result.Discount)
	require.Equal(t, arg.TotalAmount, result.TotalAmount)
	require.Equal(t, arg.PaymentInfo, result.PaymentInfo)
	require.Equal(t, "USD", result.BillingCurrency)
	require.Equal(t, arg.Note, result.Note)
	require.NotZero(t, result.CreatedAt)

	require.Len(t, result.LineItems, n)
	for i := 0; i < n; i++ {
		require.NotZero(t, result.LineItems[i].ID)
		require.Equal(t, result.InvoiceNumber, result.LineItems[i].InvoiceNumber)
		require.Equal(t, items[i].Description, result.LineItems[i].Description)
		require.Equal(t, items[i].Quantity, result.LineItems[i].Quantity)
		require.Equal(t, items[i].UnitPrice, result.LineItems[i].UnitPrice)
		require.Equal(t, items[i].TotalPrice, result.LineItems[i].TotalPrice)
	}
}
