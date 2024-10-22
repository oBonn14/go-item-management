package customer

import (
	"Go-CRUD/app/util/pagination"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type CustomerService interface {
	CreateCustomer(customer CreateCustomerRequest) (Customer, error)
	GetAllCUstomer(page *pagination.Request) ([]Customer, int64, error)
	UpdateCustomer(id int, req UpdateCustomerRequest) (Customer, error)
	DeleteCustomer(id int) error
}

type service struct {
	repo CustomerRepository
}

func NewService(repo CustomerRepository) *service {
	return &service{repo}
}

func (s service) CreateCustomer(req CreateCustomerRequest) (Customer, error) {

	tx := s.repo.DB().Begin()
	if tx.Error != nil {
		log.Error("Failed to begin transaction", tx.Error)
		return Customer{}, tx.Error
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			log.Error("Recovered from panic", tx.Error)
		}
	}()

	existingCustomer, err := s.repo.FindByName(req.CustomerName)
	if err != nil {
		return Customer{}, err
	}

	if existingCustomer != nil {
		return Customer{}, fiber.NewError(fiber.StatusConflict, "Customer "+req.CustomerName+" already exists")
	}

	customer := Customer{
		CustomerName: req.CustomerName,
		Age:          req.CustomerAge,
	}

	customer, err = s.repo.Save(tx, &customer)
	if err != nil {
		tx.Rollback()
		log.Error("Failed to data customer: ", err)
		return Customer{}, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Failed to commit transaction: ", err)
		return Customer{}, err
	}

	return customer, nil
}

func (s service) UpdateCustomer(id int, req UpdateCustomerRequest) (Customer, error) {
	customer, err := s.repo.FindById(id)
	if err != nil {
		log.Error(err)
		return Customer{}, err
	}

	customer.CustomerName = req.CustomerName
	customer.Age = req.CustomerAge

	updateCustomer, err := s.repo.Update(customer)
	if err != nil {
		log.Error(err)
		return Customer{}, err
	}

	return updateCustomer, nil
}

func (s service) GetAllCUstomer(page *pagination.Request) ([]Customer, int64, error) {
	customers, totalRecord, err := s.repo.FindAllCustomer(page)

	if err != nil {
		log.Error("Failed to get all customers: ", err)
		return nil, 0, err
	}

	return customers, totalRecord, nil
}

func (s service) DeleteCustomer(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
