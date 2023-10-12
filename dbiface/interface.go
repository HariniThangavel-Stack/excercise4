package dbiface

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionAPI interface {
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)

	DeleteOne(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)

	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (*mongo.Cursor, error)

	FindOneAndUpdate(ctx context.Context, filter interface{},
		update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult

	ReplaceOne(ctx context.Context, filter interface{},
		replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
}
