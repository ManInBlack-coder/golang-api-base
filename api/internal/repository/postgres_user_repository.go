package repository

import (
	"database/sql"
	"fmt"
	"time"

	squirrel "github.com/Masterminds/squirrel"
	"golang-api/internal/models"
)

// PostgresUserRepository implements UserRepository for PostgreSQL
type PostgresUserRepository struct {
	db      *sql.DB
	builder squirrel.StatementBuilderType
}

// NewPostgresUserRepository creates a new PostgresUserRepository instance
func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

// GetAll returns all users from the database
func (r *PostgresUserRepository) GetAll() ([]models.User, error) {
	query := r.builder.Select("id", "name", "email", "created_at", "updated_at").
		From("users").
		OrderBy("id")

	queryStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := r.db.Query(queryStr, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// GetByID returns a user by ID from the database
func (r *PostgresUserRepository) GetByID(id int) (models.User, error) {
	query := r.builder.Select("id", "name", "email", "created_at", "updated_at").
		From("users").
		Where(squirrel.Eq{"id": id})

	queryStr, args, err := query.ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to build query: %w", err)
	}

	var user models.User
	err = r.db.QueryRow(queryStr, args...).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found")
		}
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return user, nil
}

// Create inserts a new user into the database
func (r *PostgresUserRepository) Create(user models.User) (models.User, error) {
	now := time.Now()
	query := r.builder.Insert("users").
		Columns("name", "email", "created_at", "updated_at").
		Values(user.Name, user.Email, now, now).
		Suffix("RETURNING id, created_at, updated_at")

	queryStr, args, err := query.ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to build query: %w", err)
	}

	err = r.db.QueryRow(queryStr, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return user, nil
}

// Update updates an existing user in the database
func (r *PostgresUserRepository) Update(id int, user models.User) (models.User, error) {
	query := r.builder.Update("users").
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": id})

	if user.Name != "" {
		query = query.Set("name", user.Name)
	}
	if user.Email != "" {
		query = query.Set("email", user.Email)
	}

	query = query.Suffix("RETURNING id, name, email, created_at, updated_at")

	queryStr, args, err := query.ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to build query: %w", err)
	}

	var updatedUser models.User
	err = r.db.QueryRow(queryStr, args...).Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found")
		}
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return updatedUser, nil
}

// Delete removes a user from the database
func (r *PostgresUserRepository) Delete(id int) error {
	query := r.builder.Delete("users").
		Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	result, err := r.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
