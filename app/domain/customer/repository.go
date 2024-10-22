package customer

import (
	"Go-CRUD/app/util/pagination"
	"errors"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(tx *gorm.DB, customer *Customer) (Customer, error)
	FindByName(name string) (*Customer, error)
	FindById(id int) (*Customer, error)
	FindAllCustomer(page *pagination.Request) (Customers, int64, error)
	Update(customer *Customer) (Customer, error)
	Delete(id int) error
	DB() *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func (r repository) DB() *gorm.DB {
	return r.db
}

func NewCustomerRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) Save(tx *gorm.DB, customer *Customer) (Customer, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.Create(&customer).Error; err != nil {
		return Customer{}, err
	}

	return *customer, nil
}

func (r repository) Update(customer *Customer) (Customer, error) {
	if err := r.db.Save(&customer).Error; err != nil {
		return Customer{}, err
	}
	return *customer, nil
}

func (r repository) FindByName(name string) (*Customer, error) {
	var customer Customer
	if err := r.db.Where("customer_name = ?", name).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &customer, nil
}

func (r repository) FindAllCustomer(page *pagination.Request) (Customers, int64, error) {
	var customer Customers
	var totalPages int64

	query := r.db.Model(&Customer{})

	query.Count(&totalPages)

	if err := query.Scopes(pagination.Paginate(page)).Find(&customer).Error; err != nil {
		return nil, 0, err
	}
	return customer, totalPages, nil
}

func (r repository) FindById(id int) (*Customer, error) {
	var customer Customer
	if err := r.db.Where("id_customer = ?", id).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r repository) Delete(id int) error {
	if err := r.db.Where("id_customer = ?", id).Delete(&Customer{}).Error; err != nil {
		return err
	}
	return nil
}
