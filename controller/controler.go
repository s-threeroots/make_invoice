package controller

import (
	"fmt"
	"make_invoice/model"
	"make_invoice/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddInvoice(c *gin.Context) {

	invoice := &model.Invoice{}

	binderr := c.Bind(invoice)
	if binderr != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	err := service.InsertInvoice(invoice)

	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/invoice/%d", invoice.ID))
}

func DeleteInvoice(c *gin.Context) {

	paramid := c.Param("id")

	err := service.DeleteInvoice(paramid)

	if err != nil {
		if err.Error() == "record not found" {
			c.String(http.StatusNotFound, "Not Found")
		} else {
			c.String(http.StatusInternalServerError, "Server Error")
		}
		println(err)
		return
	}

	c.Redirect(http.StatusFound, "/invoice")
}

func UpadateInvoice(c *gin.Context) {

	paramid := c.Param("id")

	invoice := &model.Invoice{}

	binderr := c.Bind(invoice)
	if binderr != nil {
		fmt.Println(binderr.Error())
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	if paramid != strconv.FormatUint(uint64(invoice.ID), 10) {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	err := service.UpdateInvoice(invoice)

	if err != nil {
		if err.Error() == "record not found" {
			c.String(http.StatusNotFound, "Not Found")
		} else {
			c.String(http.StatusInternalServerError, "Server Error")
		}
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/invoice/%d", invoice.ID))
}

func AddItem(c *gin.Context) {

	item := &model.Item{}

	binderr := c.Bind(item)

	if binderr != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	err := service.InsertItem(item)

	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   item,
	})
}

func InvoiceListAll(c *gin.Context) {
	invoices, err := service.ListInvoice()

	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    invoices,
	})
}

func OneInvoiceById(c *gin.Context) {

	id := c.Param("id")

	invoice, err := service.SelectInvoiceById(id)

	if err != nil {
		if err.Error() == "record not found" {
			c.String(http.StatusNotFound, "Not Found")
		} else {
			c.String(http.StatusInternalServerError, "Server Error")
		}
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    invoice,
	})
}
