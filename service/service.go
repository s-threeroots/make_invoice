package service

import (
	"fmt"
	"make_invoice/dbutil"
	"make_invoice/model"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InsertInvoice(invoice *model.Invoice) error {

	err := dbutil.Insert(invoice)
	return err
}

func UpdateInvoice(invoice *model.Invoice) error {

	db := dbutil.Open()
	defer db.Close()

	details := invoice.Detail

	total := 0

	for i, detail := range details {

		subtotal := detail.Count * detail.Item.Price

		details[i].SubTotal = subtotal

		total += subtotal

		fmt.Println(i, detail)

	}

	invoice.Total = total

	fmt.Println(invoice)

	result := db.Save(&invoice)
	err := result.Error

	return err

}

func InsertItem(item *model.Item) error {

	err := dbutil.Insert(item)
	return err
}

func ListInvoice() (*[]model.Invoice, error) {

	invoices := []model.Invoice{}

	db := dbutil.Open()
	defer db.Close()
	result := db.Debug().Preload("Detail").Find(&invoices)
	err := result.Error

	if err != nil {
		return &invoices, err
	}

	for i, invoice := range invoices {

		setDetailStructure(&invoices[i])
		fmt.Print(invoices[i], invoice)

	}

	return &invoices, err
}

func setDetailStructure(invoice *model.Invoice) error {

	db := dbutil.Open()
	defer db.Close()

	details := invoice.Detail

	for i, detail := range details {

		result := db.Debug().First(&details[i].Item, detail.ItemID)
		err := result.Error

		fmt.Println(i, detail)
		if err != nil {
			return err
		}
	}

	return nil

}

func SelectInvoiceById(id string) (*model.Invoice, error) {

	ret := model.Invoice{}
	details := []model.Detail{}

	db := dbutil.Open()
	defer db.Close()
	result := db.Debug().First(&ret, id).Association("Detail").Find(&details)
	err := result.Error

	if err != nil {
		return &ret, err
	}

	ret.Detail = details

	err = setDetailStructure(&ret)

	if err != nil {
		return &ret, err
	}

	return &ret, err
}

func DeleteInvoice(id string) error {

	invoice := model.Invoice{}

	db := dbutil.Open()
	defer db.Close()

	err := db.Delete(invoice, id).Error
	return err
}
