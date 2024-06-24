package main

import (
	controller "Facopi-Vol2/Controller"
	connection "Facopi-Vol2/Models/Connection"

	"github.com/gin-gonic/gin"
)

func main() {

	port := ":1213"
	r := gin.Default()
	connection.ConnectDatabase()

	 r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            return
        }
        c.Next()
    })



	//###BEGIN WEB API Keterangan
	// Get data
	r.GET("/api/DataMahasiswa/GetDataFacopi", controller.GetDataFacopiByID)

	//Insert data
	r.POST("/api/DataMahasiswa/InsertDataFacopi", controller.InsertDataFacopi)

	// Update data
	r.PUT("/api/DataMahasiswa/UpdateDataFacopi", controller.UpdateDataFacopi)

	//Delete data
	r.DELETE("/api/DataMahasiswa/DeleteDataFacopi", controller.DeleteDataFacopi)
	//###END WEB API Keterangan

	//###BEGIN WEB API Keterangan
	// Get data
	r.GET("/api/DataDataFacopi/GetKeuanganFacopi", controller.GetKeuanganFacopiByID)

	//Insert data
	r.POST("/api/DataFacopi/InsertKeuanganFacopi", controller.InsertKeuanganFacopi)

	// Update data
	r.PUT("/api/DataFacopi/UpdateKeuanganFacopi", controller.UpdateKeuanganFacopi)

	//Delete data
	r.DELETE("/api/DataFacopi/DeleteKeuanganFacopi", controller.DeleteKeuanganFacopi)
	//###END WEB API Keterangan


	r.Run(port)
}