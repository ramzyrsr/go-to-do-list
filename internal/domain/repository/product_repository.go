package repository

import (
	"database/sql"
	"to-do-list/internal/domain/models"
)

type ProductRepository interface {
	Create(product *models.Product) (*models.Product, error)
	GetByID(id int) (*models.Product, error)
	Update(product *models.Product) (*models.Product, error)
	Delete(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) (*models.Product, error) {
	query := "INSERT INTO products(name, price) VALUES($1, $2) RETURNING id"
	err := r.db.QueryRow(query, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) GetByID(id int) (*models.Product, error) {
	product := &models.Product{}
	query := "SELECT applicant_id, first_name, last_name FROM recruitment_schema.trx_kandidat_data_diri WHERE verifikasi_kandidat_id = $1"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&product.ApplicantId, &product.FirstName, &product.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Update(product *models.Product) (*models.Product, error) {
	query := "UPDATE products SET name = $1, price = $2 WHERE id = $3"
	_, err := r.db.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
