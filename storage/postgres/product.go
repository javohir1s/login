package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/asadbekGo/market_system/models"
	"github.com/asadbekGo/market_system/pkg/helpers"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(ctx context.Context, req *models.CreateProduct) (*models.Product, error) {

	var (
		productId = uuid.New().String()
		query     = `
			INSERT INTO "product"(
				"id",
				"name",
				"barcode",
				"price",
				"image_url",
				"category_id",
				"updated_at"
			) VALUES ($1, $2, $3, $4, $5, $6, NOW())`
	)

	_, err := r.db.Exec(ctx,
		query,
		productId,
		req.Name,
		req.Barcode,
		req.Price,
		req.ImageUrl,
		helpers.NewNullString(req.CategoryId),
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, &models.ProductPrimaryKey{Id: productId})
}

func (r *productRepo) GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {

	var (
		query = `
			SELECT
				"id",
				"name",
				"barcode",
				"price",
				"image_url",
				"category_id",
				"updated_at",
				"created_at"
			FROM "product"
			WHERE "id" = $1
		`
	)

	var (
		id          sql.NullString
		name        sql.NullString
		barcode     sql.NullString
		price       sql.NullFloat64
		image_url   sql.NullString
		category_id sql.NullString
		updated_at  sql.NullString
		created_at  sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&barcode,
		&price,
		&image_url,
		&category_id,
		&updated_at,
		&created_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Product{
		Id:         id.String,
		Name:       name.String,
		Barcode:    barcode.String,
		Price:      price.Float64,
		ImageUrl:   image_url.String,
		CategoryId: category_id.String,
		UpdatedAt:  updated_at.String,
		CreatedAt:  created_at.String,
	}, nil
}

func (r *productRepo) GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error) {
	var (
		resp   models.GetListProductResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY created_at DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.Search) > 0 {
		where += " AND name ILIKE" + " '%" + req.Search + "%'"
	}

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"name",
			"barcode",
			"price",
			"image_url",
			"category_id",
			"updated_at",
			"created_at"
		FROM "product"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			name        sql.NullString
			barcode     sql.NullString
			price       sql.NullFloat64
			image_url   sql.NullString
			category_id sql.NullString
			updated_at  sql.NullString
			created_at  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&barcode,
			&price,
			&image_url,
			&category_id,
			&updated_at,
			&created_at,
		)
		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &models.Product{
			Id:         id.String,
			Name:       name.String,
			Barcode:    barcode.String,
			Price:      price.Float64,
			ImageUrl:   image_url.String,
			CategoryId: category_id.String,
			UpdatedAt:  updated_at.String,
			CreatedAt:  created_at.String,
		})
	}

	return &resp, nil
}

func (r *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {

	query := `
		UPDATE product
			SET
				name = $2,
				barcode = $3,
				price = $4,
				image_url = $5,
				category_id = $6,
		WHERE id = $1
	`
	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.Name,
		req.Barcode,
		req.Price,
		req.ImageUrl,
		helpers.NewNullString(req.CategoryId),
	)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (r *productRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) error {
	_, err := r.db.Exec(ctx, "DELETE FROM product WHERE id = $1", req.Id)
	return err
}
