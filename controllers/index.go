package rpcService

import (
	"context"
	"fmt"
	"inventory/interfaces"
	pb "inventory/proto"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type RPCServer struct {
	mu sync.Mutex
	pb.UnimplementedInventoryServiceServer
}

var (
	InventoryService interfaces.Inventory
	Mcoll            *mongo.Collection
)

func (s *RPCServer) GetAllItems(ctx context.Context, req *pb.Empty) (*pb.AllInventoryItems, error) {
	fmt.Println("GetAllItems")
	s.mu.Lock()
	defer s.mu.Unlock()
	res, err := InventoryService.GetAllItems()
	if err != nil {
		return nil, err
	}
	inventory := []*pb.InventoryItem{}
	for i := 0; i < len(res); i++ {
		ivs := []*pb.InventorySKU{}
		for j := 0; j < len(res[i].Skus); j++ {
			ivt := &pb.InventorySKU{
				Sku: res[i].Skus[j].Sku,
				Price: &pb.Price{
					Base:     res[i].Skus[j].Price.Base,
					Currency: res[i].Skus[j].Price.Currency,
					Discount: res[i].Skus[j].Price.Discount,
				},
				Quantity: res[i].Skus[j].Quantity,
				Options: &pb.Options{
					Size: &pb.Size{
						H: res[i].Skus[j].Options.Size.H,
						L: res[i].Skus[j].Options.Size.L,
						W: res[i].Skus[j].Options.Size.W,
					},
					Features: res[i].Skus[j].Options.Features,
					Colors:   res[i].Skus[j].Options.Colors,
					Ruling:   res[i].Skus[j].Options.Ruling,
					Image:    res[i].Skus[j].Options.Image,
				},
			}
			ivs = append(ivs, ivt)
		}
		iv := &pb.InventoryItem{
			Id:         res[i].ID,
			Item:       res[i].Item,
			Features:   res[i].Features,
			Categories: res[i].Categories,
			Skus:       ivs,
		}
		inventory = append(inventory, iv)

	}
	fmt.Println(inventory)

	return &pb.AllInventoryItems{
		Items: inventory,
	}, nil
}

func (s *RPCServer) GetInventoryItemByItemName(ctx context.Context, req *pb.ItemName) (*pb.InventoryItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, err := InventoryService.GetInventoryItemByItemName(req.ItemName)
	if err != nil {
		return nil, err
	}
	ivs := []*pb.InventorySKU{}
	for j := 0; j < len(res.Skus); j++ {
		ivt := &pb.InventorySKU{
			Sku: res.Skus[j].Sku,
			Price: &pb.Price{
				Base:     res.Skus[j].Price.Base,
				Currency: res.Skus[j].Price.Currency,
				Discount: res.Skus[j].Price.Discount,
			},
			Quantity: res.Skus[j].Quantity,
			Options: &pb.Options{
				Size: &pb.Size{
					H: res.Skus[j].Options.Size.H,
					L: res.Skus[j].Options.Size.L,
					W: res.Skus[j].Options.Size.W,
				},
				Features: res.Skus[j].Options.Features,
				Colors:   res.Skus[j].Options.Colors,
				Ruling:   res.Skus[j].Options.Ruling,
				Image:    res.Skus[j].Options.Image,
			},
		}
		ivs = append(ivs, ivt)
	}
	inventory := &pb.InventoryItem{
		Id:         res.ID,
		Item:       res.Item,
		Features:   res.Features,
		Categories: res.Categories,
		Skus:       ivs,
	}
	return inventory, nil
}


