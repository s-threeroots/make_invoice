package main

import (
	"make_invoice/controller"
	"make_invoice/dbutil"
	"make_invoice/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	dbutil.InitDb()
	engine := gin.Default()

	engine.Use(middleware.RecordUaAndTime)

	invoiceEngine := engine.Group("/invoice")
	{

		{
			invoiceEngine.GET("/", controller.InvoiceListAll)
			invoiceEngine.GET("/:id", controller.OneInvoiceById)
			invoiceEngine.POST("/:id", controller.UpadateInvoice)
			invoiceEngine.DELETE("/:id", controller.DeleteInvoice)
			invoiceEngine.POST("/add", controller.AddInvoice)

		}
	}

	engine.Run(":3000")
}
