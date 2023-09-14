package main

import (
	grpcclient "inventory/client/grpcClient"
	"inventory/client/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	_, conn := grpcclient.GetGrpcClientInstance()
	defer conn.Close()

	r := gin.Default()
	routes.AppRoutes(r)
	r.Run(":8080")
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	h "inventory/proto"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	r := gin.Default()
// 	conn, err := grpc.Dial("localhost:2023", grpc.WithInsecure())
// 	if err != nil {
// 		fmt.Println("Failed to connect: ", err)
// 	}
// 	defer conn.Close()
// 	client := h.NewInventoryServiceClient(conn)
// 	r.GET("/getitems", func(c *gin.Context) {
// 		response, err := client.GetAllItems(context.Background(), &h.Empty{})
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}
// 		fmt.Println("Response: ", response)
// 		c.JSON(http.StatusOK, gin.H{"value": response})
// 	})

// 	r.POST("/updateitems", func(c *gin.Context) {
// 		var request h.ItemToDelete
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		response, err := client.UpdateInventory(context.Background(), &h.ItemToDelete{
// 			Item:     request.Item,
// 			Sku:      request.Sku,
// 			Quantity: request.Quantity,
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}
// 		fmt.Println("Response: ", response)
// 		c.JSON(http.StatusOK, gin.H{"value": response})
// 	})

// 	r.POST("/getitem", func(c *gin.Context) {mongodb+srv://vigneshk:1234@banking.obgwcv6.mongodb.net/?retryWrites=true&w=majority
// 		var request h.ItemName
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		response, err := client.GetInventoryItemByItemName(context.Background(), &h.ItemName{
// 			ItemName: request.ItemName,
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}
// 		fmt.Println("Response: ", response)
// 		c.JSON(http.StatusOK, gin.H{"value": response})
// 	})
// 	r.POST("/create", func(c *gin.Context) {
// 		var request []*h.InventoryItem
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		response, err := client.CreateInventory(context.Background(), &h.AllInventoryItems{
// 			Items: request,
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}
// 		fmt.Println("Response: ", response)
// 		c.JSON(http.StatusOK, gin.H{"value": response})
// 	})
// 	r.Run(":5000")
// }
