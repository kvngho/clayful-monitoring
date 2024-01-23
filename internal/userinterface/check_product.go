package userinterface

import (
	"context"

	"github.com/kvngho/clayful-monitoring/internal/domain/db"
	"github.com/kvngho/clayful-monitoring/internal/domain/ent/product"
	"github.com/kvngho/clayful-monitoring/internal/infra"
	"github.com/kvngho/clayful-monitoring/internal/pkg/clayful"
)

type Checker struct {
	client infra.EntDB
	clayful infra.Clayful
}

func NewChecker(db infra.EntDB, clayful infra.Clayful) *Checker {
	return &Checker{
		client: db,
		clayful: clayful,
	}
}

func (c Checker) FindProductDeepingAndClayful(ctx context.Context) ([]db.Product, error) {

	queryResult, err := c.client.DB.Product.Query().Where(product.ClayfulIDNEQ("")).All(context.Background())
	result := make([]db.Product, len(queryResult))
	if err != nil {
		return nil, err
	}
	for _, p := range queryResult {
		result = append(result, db.Product{
			ID: p.ID,
			Name: p.Name,
			Clayful_id: p.ClayfulID,
			Clayful_options: p.ClayfulOptions,
			Price_deeping: p.PriceDeeping,
		})
	}

	return result, nil
}

func (c Checker) FindProductFromClayful(ctx context.Context, productID string) *clayful.ClayfulProduct {
	return c.clayful.GetProduct(ctx, productID)
}