package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Connection struct {
	conn *mongo.Client
}

type Database struct {
	db *mongo.Database
}

type Collection struct {
	delegate *mongo.Collection
}

//初始化
func New(uri string) (connector *Connection, err error) {

	client, err := setConnect(uri)
	if err != nil {
		fmt.Println(err)
		return connector, err
	}

	connector = &Connection{
		conn: client,
	}
	return connector, err
}

// 连接设置
func setConnect(uri string) (client *mongo.Client, err error) {

	client, err = mongo.Connect(getContext(), options.Client().ApplyURI(uri)) // 连接池

	if err == nil {
		err = client.Ping(getContext(), readpref.Primary())
	}
	return client, err
}

func getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}
