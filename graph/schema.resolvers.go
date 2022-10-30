package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/database"
	"example/graph/generated"
	"example/graph/model"
	"fmt"
)

// CreateNewCardTierScore is the resolver for the createNewCardTierScore field.
func (r *mutationResolver) CreateNewCardTierScore(ctx context.Context, input model.NewCardTier) (*model.CardTier, error) {
	panic(fmt.Errorf("not implemented: CreateNewCardTierScore - createNewCardTierScore"))
}

// CardTier is the resolver for the cardTier field.
func (r *queryResolver) CardTier(ctx context.Context, id string) (*model.CardTier, error) {
	return db.FindCardTierScoreById(id), nil
}

// CardTiers is the resolver for the cardTiers field.
func (r *queryResolver) CardTiers(ctx context.Context) ([]*model.CardTier, error) {
	panic(fmt.Errorf("not implemented: CardTiers - cardTiers"))
}

// CardTierByName is the resolver for the cardTierByName field.
func (r *queryResolver) CardTierByName(ctx context.Context, name string) (*model.CardTier, error) {
	return db.FindCardTierScoreByName(name), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) CardTierByID(ctx context.Context, id string) (*model.CardTier, error) {
	panic(fmt.Errorf("not implemented: CardTierByID - cardTierById"))
}

var db = database.Connect("mongodb://localhost:27017/")
