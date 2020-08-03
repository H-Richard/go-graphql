package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/H-Richard/go-graphql/database"
	"github.com/H-Richard/go-graphql/graph/generated"
	"github.com/H-Richard/go-graphql/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
	return db.Save(input), nil
}

func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
	return db.FindByID(id), nil
}

func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	return db.All(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
