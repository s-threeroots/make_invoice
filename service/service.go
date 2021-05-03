package service

import (
	"database/sql"
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

	ret := []model.Invoice{}

	db := dbutil.Open()
	defer db.Close()
	rows, err := db.Debug().
		Table("invoices").
		Select("invoices.*, details.*, items.*").
		Joins("left join details on details.invoice_id = invoices.id").
		Joins("inner join items on details.item_id = items.id").
		Where("invoices.deleted_at is null").
		Order("invoices.id").Rows()

	if err != nil {
		return &ret, err
	}

	defer rows.Close()

	return getInvoiceFromRows(rows)

}

func getInvoiceFromRows(rows *sql.Rows) (*[]model.Invoice, error) {

	ret := []model.Invoice{}
	var err error = nil

	for rows.Next() {

		invoice := model.Invoice{}
		detail := model.Detail{}
		item := model.Item{}

		err = rows.Scan(
			&invoice.ID,
			&invoice.CreatedAt,
			&invoice.UpdatedAt,
			&invoice.DeletedAt,
			&invoice.Name,
			&invoice.Name,
			&invoice.Total,
			&detail.InvoiceID,
			&detail.ItemID,
			&detail.Count,
			&detail.SubTotal,
			&item.ID,
			&item.Name,
			&item.Price,
			&item.Unit,
			&item.Description)

		if err != nil {
			return &ret, err
		}

		if len(ret) == 0 || ret[len(ret)-1].ID != invoice.ID {
			ret = append(ret, invoice)
		}
		invoice = ret[len(ret)-1]
		detail.Item = item
		ret[len(ret)-1].Detail = append(ret[len(ret)-1].Detail, detail)
	}

	return &ret, err

}

func SelectInvoiceById(id string) (*model.Invoice, error) {

	ret := model.Invoice{}

	db := dbutil.Open()
	defer db.Close()

	rows, err := db.Debug().
		Table("invoices").
		Select("invoices.*, details.*, items.*").
		Joins("left join details on details.invoice_id = invoices.id").
		Joins("inner join items on details.item_id = items.id").
		Where("invoices.id = ?", id).Where("invoices.deleted_at is null").Rows()

	if err != nil {
		return &ret, err
	}

	defer rows.Close()

	invoices, err := getInvoiceFromRows(rows)

	if err != nil {
		return &ret, err
	}

	ret = (*invoices)[0]

	return &ret, err
}

func DeleteInvoice(id string) error {

	invoice := model.Invoice{}

	db := dbutil.Open()
	defer db.Close()

	err := db.Delete(invoice, id).Error
	return err
}
