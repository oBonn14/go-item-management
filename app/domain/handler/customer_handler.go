package handler

import (
	"Go-CRUD/app/domain/customer"
	"Go-CRUD/app/util/pagination"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerService customer.CustomerService
}

func NewCustomerHandler(customerService customer.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService}
}

func (h *CustomerHandler) AddCustomer(c *fiber.Ctx) error {
	var req customer.CreateCustomerRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	addCustomer, err := h.customerService.CreateCustomer(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Customer created successfully",
		"customer": addCustomer,
	})
}

func (h *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var req customer.UpdateCustomerRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	UpdateCustomer, err := h.customerService.UpdateCustomer(id, req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Customer updated successfully",
		"data":    UpdateCustomer,
	})
}

func (h *CustomerHandler) GetAllCustomer(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	size := c.QueryInt("size", 10)

	paginationRequest := pagination.Request{
		Page: page,
		Size: size,
	}

	customer, totalRecord, err := h.customerService.GetAllCUstomer(&paginationRequest)
	if err != nil {
		return err
	}

	pagination := pagination.New(totalRecord, page, size)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    customer,
		"total":   pagination.TotalRecords,
		"page":    pagination.Page,
		"size":    pagination.Size,
	})
}

func (h *CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.customerService.DeleteCustomer(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Customer deleted successfully",
	})
}
