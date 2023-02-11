package usecase

import (
	"context"
	"time"

	"github.com/emerfauzan/cake-store-api/app/model"
	"github.com/emerfauzan/cake-store-api/app/request"
	"github.com/emerfauzan/cake-store-api/app/response"
)

func (usecase *Usecase) GetCakes(ctx context.Context) (result []response.CakeResponse, err error) {
	cakes, err := usecase.repository.GetActiveCakes(ctx)
	if err != nil {
		return result, err
	}

	for _, cake := range cakes {
		cakeResponse := response.CakeResponse{
			Id:          cake.Id,
			Title:       cake.Title,
			Description: cake.Description,
			Rating:      cake.Rating,
			ImageUrl:    cake.ImageUrl,
		}

		result = append(result, cakeResponse)
	}
	return result, nil
}

func (usecase *Usecase) GetCakeById(ctx context.Context, id uint) (result response.CakeResponse, err error) {
	cake, err := usecase.repository.GetCakeById(ctx, id)
	if err != nil {
		return result, err
	}

	result.Id = cake.Id
	result.Title = cake.Title
	result.Description = cake.Description
	result.Rating = cake.Rating
	result.ImageUrl = cake.ImageUrl

	return result, nil
}

func (usecase *Usecase) CreateCake(ctx context.Context, request request.CreateCakeRequest) (err error) {
	cake := model.Cake{
		Title:       request.Title,
		Description: request.Description,
		Rating:      request.Rating,
		ImageUrl:    request.ImageUrl,
		CreatedBy:   request.UserId,
		CreatedAt:   time.Now(),
	}
	err = usecase.repository.CreateCake(ctx, cake)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *Usecase) UpdateCake(ctx context.Context, request request.UpdateCakeRequest, id uint) (err error) {
	cake, err := usecase.repository.GetCakeById(ctx, id)
	if err != nil {
		return err
	}

	cake.Title = request.Title
	cake.Description = request.Description
	cake.Rating = request.Rating
	cake.ImageUrl = request.ImageUrl
	cake.UpdatedBy = request.UserId
	cake.UpdatedAt = time.Now()
	err = usecase.repository.UpdateCake(ctx, cake)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *Usecase) DeleteCake(ctx context.Context, id uint, userId uint) (err error) {
	cake, err := usecase.repository.GetCakeById(ctx, id)
	if err != nil {
		return err
	}

	cake.IsDeletedFlag = true
	cake.DeletedBy = userId
	cake.DeletedAt = time.Now()
	err = usecase.repository.UpdateCake(ctx, cake)
	if err != nil {
		return err
	}

	return nil
}
