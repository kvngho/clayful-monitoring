package userinterface

import (
	"context"
	"time"

	"github.com/kvngho/clayful-monitoring/internal/domain/db"
	"github.com/kvngho/clayful-monitoring/internal/domain/ent/coupon"
	"github.com/kvngho/clayful-monitoring/internal/pkg/clayful"
)

func (c Checker) FindAvailableCouponDeeping(ctx context.Context) ([]db.Coupon, error){
	queryResult, err := c.client.DB.Coupon.Query().Where(coupon.EndDateLTE(time.Now())).All(context.Background())
	result := make([]db.Coupon, len(queryResult))
	if err != nil {
		return nil, err
	}
	for _, p := range queryResult {
		result = append(result, db.Coupon{
			ID: p.ID,
			Clayful_id: p.ClayfulID,
			EndDate: p.EndDate.String(),
			Name: p.Name,
			Active: p.Active,
		})
	}
	return result, nil
}

func (c Checker) FindCouponFromClayful(ctx context.Context, productID string) *clayful.ClayfulCoupon {
	return c.clayful.GetCoupon(ctx, productID)
}