package db 

import (
	"context"

	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	*mongo.Client
}


func Connect() (*Client , error) {
	
}