package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"github.com/rupesh40/go-car-rental/V1/database"
	"github.com/rupesh40/go-car-rental/V1/graph/model"
)

// Cars is the resolver for the cars field.
func (r *queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	var cars []*model.Car
	var db = database.GetDB()
	result := db.Find(&cars)

	if result.Error != nil {
		return cars, result.Error
	}
	return cars, nil
}

// Car is the resolver for the car field.
func (r *queryResolver) Car(ctx context.Context, id string) (*model.Car, error) {
	var db = database.GetDB()
		var car *model.Car
		result:=db.Where("id = ?", id).First(&car)
		if result.Error!= nil{
			return car, result.Error
		}
		return car, nil 
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.



// func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
// 	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
// }
// func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*model.User, error) {
// 	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
// }
// func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.DeleteUserResponse, error) {
// 	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
// }
// func (r *mutationResolver) CreateCar(ctx context.Context, input model.CreateCarInput) (*model.Car, error) {
// 	panic(fmt.Errorf("not implemented: CreateCar - createCar"))
// }
// func (r *mutationResolver) UpdateCar(ctx context.Context, id string, input model.UpdateCarInput) (*model.Car, error) {
// 	panic(fmt.Errorf("not implemented: UpdateCar - updateCar"))
// }
// func (r *mutationResolver) DeleteCar(ctx context.Context, id string) (*model.DeleteCarResponse, error) {
// 	panic(fmt.Errorf("not implemented: DeleteCar - deleteCar"))
// }
// func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
// 	panic(fmt.Errorf("not implemented: Users - users"))
// }
// func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
// 	panic(fmt.Errorf("not implemented: User - user"))
// }
// func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

