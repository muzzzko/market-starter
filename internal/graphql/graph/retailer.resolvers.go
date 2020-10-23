package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"market-starter/internal/domain/retailer/model"
	"market-starter/internal/error/grapherr"
	"market-starter/internal/graphql/graph/generated"
	model1 "market-starter/internal/graphql/graph/model"
	"market-starter/internal/operation/authorization"
	intgin "market-starter/internal/operation/gin"
)

func (r *mutationResolver) CreateRetailerEmployee(ctx context.Context, input model.NewRetailerEmployee) (*model.Employee, error) {
	employee := &model.Employee{
		FirstName:  input.FirstName,
		SecondName: input.SecondName,
		Email:      input.Email,
		Password:   input.Password,
	}

	if err := r.employeeService.CreateEmployee(ctx, employee); err != nil {
		return nil, grapherr.Serialize(err)
	}

	r.generateTokenForEmployeeAndAddToGinContext(ctx, employee)

	return employee, nil
}

func (r *mutationResolver) CreateRetailer(ctx context.Context, input *model1.NewRetailer) (*model.Retailer, error) {
	id := authorization.EmployeeSessionFromContext(ctx).EmployeeID

	employee := r.employeeService.GetEmployee(ctx, id)

	retailer := &model.Retailer{
		Name: input.Name,
	}

	err := r.employeeService.AddRetailerForRetailerEmployee(ctx, employee, retailer, "owner")
	if err != nil {
		return nil, grapherr.Serialize(err)
	}

	return retailer, nil
}

func (r *queryResolver) LoginRetailerEmployeeByEmail(ctx context.Context, email string, password string) (*model.Employee, error) {
	employee, err := r.employeeService.CheckPasswordAndReturnEmployeeByEmail(ctx, email, password)
	if err != nil {
		return nil, grapherr.Serialize(err)
	}

	r.generateTokenForEmployeeAndAddToGinContext(ctx, employee)

	return employee, nil
}

func (r *queryResolver) RetailerEmployee(ctx context.Context) (*model.Employee, error) {
	id := authorization.EmployeeSessionFromContext(ctx).EmployeeID

	return r.employeeService.GetEmployee(ctx, id), nil
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
func (r *retailerResolver) EmployeeRole(ctx context.Context, obj *model.Retailer) (*model.RetailerEmployeeRole, error) {
	panic(fmt.Errorf("not implemented"))
}

type retailerResolver struct{ *Resolver }
type retailerEmployeeResolver struct{ *Resolver }

func (r *retailerEmployeeResolver) Retailer(ctx context.Context, obj *model.Employee) (*model.Retailer, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *Resolver) generateTokenForEmployeeAndAddToGinContext(ctx context.Context, employee *model.Employee) {
	session := r.sessionManager.InitEmployeeSession(employee)
	r.sessionManager.SetEmployeeSession(ctx, session)

	claims := r.jwtService.GetEmployeeClaims(ctx, employee)
	r.jwtService.AddSessionIDClaims(claims, session.SessionID)
	token := r.jwtService.Generate(ctx, claims)

	ginctx, _ := intgin.GinContextFromContext(ctx)
	ginctx.SetCookie("auth", token, 3600, "/", "localhost", false, true)
}
