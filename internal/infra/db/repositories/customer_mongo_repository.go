package repositories

import (
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain/entities"
	"github.com/valyala/fasthttp"
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
func (r *CustomerMongoRepository) FindById(ctx *fasthttp.RequestCtx, id string) (entities.Customer, error) {

	filter := bson.M{"_id": id}
	var c entities.Customer
	startTime := time.Now()
	if err := r.customerCollection.FindOne(ctx, filter).Decode(&c); err != nil {
		if err == mongo.ErrNoDocuments {
			return entities.Customer{}, ErrCustomerNotFound
		}
	}
	endTime := time.Now()
	log.Infof("Reading from DB took %v", endTime.Sub(startTime))
	return c, nil
}
func (r *CustomerMongoRepository) Save(ctx *fasthttp.RequestCtx, c entities.Customer) error {
	startTime := time.Now()
	filter := bson.D{{Key: "_id", Value: c.Id}, {Key: "version", Value: c.Version}}
	c.Version++

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "balance", Value: c.Balance},
		{Key: "version", Value: c.Version},
		{Key: "transactions", Value: c.Transactions},
	}}}
	u, err := r.customerCollection.UpdateOne(ctx, filter, update)
	endTime := time.Now()
	log.Infof("Saving to db took %v", endTime.Sub(startTime))
	if err != nil {
		log.Fatalf("Error saving to db %v", err)
		return err
	}
	if u.ModifiedCount == 0 {
		return ErrCustomerNotUpdated
	}

	return nil
}
