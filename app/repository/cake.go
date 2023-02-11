package repository

import (
	"context"

	"github.com/emerfauzan/cake-store-api/app/model"
	"github.com/emerfauzan/cake-store-api/lib/logger"
)

func (repository *Repository) GetCakeById(ctx context.Context, id uint) (cake model.Cake, err error) {
	rows, err := repository.db.Query("select id, title, description, rating, image_url, is_deleted_flag, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at from cakes where id = (?)", id)

	if err != nil {
		logger.Error(ctx, "error get cake by id", map[string]interface{}{
			"error": err,
			"tags":  []string{"query"},
		})
		return cake, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.ImageUrl, &cake.IsDeletedFlag, &cake.CreatedBy, &cake.UpdatedBy, &cake.DeletedBy, &cake.CreatedAt, &cake.UpdatedAt, &cake.DeletedAt); err != nil {
			logger.Error(ctx, "error get cake by id", map[string]interface{}{
				"error": err,
				"tags":  []string{"query"},
			})
			return cake, err
		}
	}

	return cake, nil
}

func (repository *Repository) GetActiveCakes(ctx context.Context) (cakes []model.Cake, err error) {
	rows, err := repository.db.Query("selects id, title, description, rating, image_url, is_deleted_flag, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at from cakes where is_deleted_flag = false")

	if err != nil {
		logger.Error(ctx, "error get active cakes", map[string]interface{}{
			"error": err,
			"tags":  []string{"query"},
		})
		return cakes, err
	}
	defer rows.Close()

	for rows.Next() {
		var cake model.Cake
		if err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.ImageUrl, &cake.IsDeletedFlag, &cake.CreatedBy, &cake.UpdatedBy, &cake.DeletedBy, &cake.CreatedAt, &cake.UpdatedAt, &cake.DeletedAt); err != nil {
			logger.Error(ctx, "error get active cakes", map[string]interface{}{
				"error": err,
				"tags":  []string{"query"},
			})
			return cakes, err
		}

		cakes = append(cakes, cake)
	}

	return cakes, nil
}

func (repository *Repository) CreateCake(ctx context.Context, cake model.Cake) error {
	_, err := repository.db.Exec("INSERT INTO cakes (title, description, rating, image_url, is_deleted_flag, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at) values (?,?,?,?,?,?,?,?,?,?,?)",
		cake.Title, cake.Description, cake.Rating, cake.ImageUrl, cake.IsDeletedFlag, cake.CreatedBy,
		cake.UpdatedBy, cake.DeletedBy, cake.CreatedAt, cake.UpdatedAt, cake.DeletedAt)

	if err != nil {
		logger.Error(ctx, "error create cake", map[string]interface{}{
			"error": err,
			"tags":  []string{"query"},
		})
		return err
	}
	return nil
}

func (repository *Repository) UpdateCake(ctx context.Context, cake model.Cake) error {
	_, err := repository.db.Exec("UPDATE cakes set title=?, description=?, rating=?, image_url=?, is_deleted_flag=?, created_by=?, updated_by=?, deleted_by=?, created_at=?, updated_at=?, deleted_at=? WHERE id=?",
		cake.Title, cake.Description, cake.Rating, cake.ImageUrl, cake.IsDeletedFlag, cake.CreatedBy,
		cake.UpdatedBy, cake.DeletedBy, cake.CreatedAt, cake.UpdatedAt, cake.DeletedAt, cake.Id)

	if err != nil {
		logger.Error(ctx, "error update cake", map[string]interface{}{
			"error": err,
			"tags":  []string{"query"},
		})
		return err
	}
	return nil
}
