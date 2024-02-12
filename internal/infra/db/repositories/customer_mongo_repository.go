package repositories

import (
	"context"

	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerMongoRepository struct {
	customerCollection mongo.Collection
}

func NewCustomerMongoRepository(client mongo.Client) *CustomerMongoRepository {
	collection := client.Database("wallet").Collection("customer")
	return &CustomerMongoRepository{
		customerCollection: *collection,
	}
}
func (r *CustomerMongoRepository) FindById(id int) (entities.Customer, error) {
	filter := bson.D{{Key: "balance", Value: 0}}
	c := entities.Customer{}
	if err := r.customerCollection.FindOne(context.TODO(), filter).Decode(c); err != nil {
		return entities.Customer{}, err
	}
	return c, nil
}
func (r *CustomerMongoRepository) Save(c entities.Customer) error {
	filter := bson.D{{Key: "_id", Value: c.Id}}
	_, err := r.customerCollection.ReplaceOne(context.TODO(), filter, c)
	return err
}
