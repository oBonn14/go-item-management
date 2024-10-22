package customer

type CreateCustomerRequest struct {
	CustomerName string `json:"name" validate:"required"`
	CustomerAge  int    `json:"Age" validate:"required"`
}

type UpdateCustomerRequest struct {
	CustomerName string `json:"name" validate:"required"`
	CustomerAge  int    `json:"Age" validate:"required"`
}
