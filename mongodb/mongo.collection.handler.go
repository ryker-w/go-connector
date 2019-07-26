package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (collection Collection) Find(query map[string]interface{}, res interface{}) (err error) {
	s := collection.delegate.FindOne(getContext(), query)
	err = s.Decode(res)
	return err
}

func (collection Collection) Update(query map[string]interface{}, change map[string]interface{}) (err error) {
	s := collection.delegate.FindOneAndUpdate(getContext(), query, bson.M{"$set": change})
	err = s.Err()
	return err
}

func (collection Collection) Replace(query map[string]interface{}, change interface{}) (err error) {
	s := collection.delegate.FindOneAndReplace(getContext(), query, change)
	err = s.Err()
	return err
}

func (collection Collection) Insert(data interface{}) (id string, err error) {
	res, err := collection.delegate.InsertOne(getContext(), data)
	if err != nil {
		fmt.Println(err)
		return id, err
	}
	id = res.InsertedID.(primitive.ObjectID).String()
	return id, err
}