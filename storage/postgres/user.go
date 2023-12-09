package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, req *models.CreateUser) (*models.User, error) {

	var (
		userId = uuid.New().String()
		query  = `
			INSERT INTO "user"(
				"id",
				"name",
				"login",
				"password",
				"created_at",
				"updated_at"
			) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	)

	_, err := r.db.Exec(ctx,
		query,
		userId,
		req.Name,
		req.Login,
		req.Password,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, &models.UserPrimaryKey{Id: userId})
}

func (r *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query = `
			SELECT
				"id",
				"name",
				"login",
				"password",
				"created_at",
				"updated_at"
			FROM "user"
			WHERE "id" = $1
		`
	)

	var (
		id        sql.NullString
		name      sql.NullString
		login     sql.NullString
		password  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&login,
		&password,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        id.String,
		Name:      name.String,
		Login:     login.String,
		Password:  password.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {
	var (
		resp   models.GetListUserResponse
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
			"login",
			"password",
			"created_at",
			"updated_at"
		FROM "user"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			name      sql.NullString
			login     sql.NullString
			password  sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&login,
			&password,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:        id.String,
			Name:      name.String,
			Login:     login.String,
			Password:  password.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return &resp, nil
}

func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
	query := `
		UPDATE "user"
			SET
				"name" = $2,
				"login" = $3,
				"password" = $4,
				"updated_at" = NOW()
		WHERE "id" = $1
	`
	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.Name,
		req.Login,
		req.Password,
	)
	if err != nil {
		return 0, err
	}
	return rowsAffected.RowsAffected(), nil
}

func (r *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {
	query := `DELETE FROM "user" WHERE "id" = $1`
	_, err := r.db.Exec(ctx, query, req.Id)
	return err
}
func (r *userRepo) GetByLoginAndPassword(ctx context.Context, req *models.LoginRequest) (*models.User, error) {
	query := `
		SELECT
			"id",
			"name",
			"login",
			"password",
			"created_at",
			"updated_at"
		FROM "user"
		WHERE "login" = $1 AND "password" = $2
	`

	var (
		id        sql.NullString
		name      sql.NullString
		login     sql.NullString
		password  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Login, req.Password).Scan(
		&id,
		&name,
		&login,
		&password,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err != nil {
			log.Println("user not found")
			return nil, errors.New("user not found")

		}
		return nil, err
	}

	return &models.User{
		Id:        id.String,
		Name:      name.String,
		Login:     login.String,
		Password:  password.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *userRepo) Refresh(ctx context.Context, login string) error {
	apiKeyExpiration := config.ApiKeyExpiredAt

	expirationTime := time.Now().Add(apiKeyExpiration)

	query := `
		UPDATE "user"
		SET
			"expired_at" = $2
		WHERE "login" = $1
	`

	_, err := r.db.Exec(ctx, query, login, expirationTime)
	if err != nil {
		return err
	}

	return nil
}
