package controller

import (
	"context"
	"fmt"
	grpcclient "inventory/client/grpcClient"
	h "inventory/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerGetAll(c *gin.Context) {

	grpcClient, _ := grpcclient.GetGrpcClientInstance()

	response, err := grpcClient.GetAllItems(context.Background(), &h.Empty{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println("Response: ", response)
	c.JSON(http.StatusOK, gin.H{"value": response})
}

func HandlerUpdateItems(c *gin.Context) {

	grpcClient, _ := grpcclient.GetGrpcClientInstance()

	var request h.ItemToDelete
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := grpcClient.UpdateInventory(context.Background(), &h.ItemToDelete{
		Item:     request.Item,
		Sku:      request.Sku,
		Quantity: request.Quantity,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println("Response: ", response)
	c.JSON(http.StatusOK, gin.H{"value": response})

}

func HandlerGetItem(c *gin.Context) {
	grpcClient, _ := grpcclient.GetGrpcClientInstance()

	var request h.ItemName
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := grpcClient.GetInventoryItemByItemName(context.Background(), &h.ItemName{
		ItemName: request.ItemName,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println("Response: ", response)
	c.JSON(http.StatusOK, gin.H{"value": response})
}

func HandlerCreate(c *gin.Context) {
	grpcClient, _ := grpcclient.GetGrpcClientInstance()

	var request []*h.InventoryItem
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := grpcClient.CreateInventory(context.Background(), &h.AllInventoryItems{
		Items: request,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println("Response: ", response)
	c.JSON(http.StatusOK, gin.H{"value": response})

}

func HandlerAddItems(c *gin.Context) {
	grpcClient, _ := grpcclient.GetGrpcClientInstance()

	var request []*h.InventorySKU
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := grpcClient.AddItems(context.Background(), &h.AllInventorySKUItems{
		Items: request,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println("Response: ", response)
	c.JSON(http.StatusOK, gin.H{"value": response})

}
