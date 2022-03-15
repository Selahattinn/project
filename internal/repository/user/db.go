package user

import (
	"context"
	"fmt"

	model "github.com/Selahattinn/bitaksi/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

const (
	dataBaseName   = "bitaksi"
	collectionName = "users"
)

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	collection := client.Database(dataBaseName).Collection(collectionName)
	return &MongoRepository{
		collection: collection,
	}

}

func (r *MongoRepository) GetUser(email string) (*model.User, error) {
	u := &model.User{}
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(u)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *MongoRepository) CreateUser(user *model.User) (*model.User, error) {
	result, err := r.collection.InsertOne(context.TODO(), bson.M{"email": user.Email})
	if err != nil {
		return nil, err
	}

	fmt.Println(result.InsertedID)
	// convert result to model.User

	return user, nil

}
