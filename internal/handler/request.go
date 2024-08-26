package handler

type CreateCustomerRequest struct {
	Name string `json:"name" validate:"required"`
	Age  uint   `json:"age" validate:"required,min=1,max=200"`
}

type UpdateCustomerRequest struct {
	ID   uint   `param:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Age  uint   `json:"age" validate:"required,min=1,max=200"`
}

type GetCustomerByIDRequest struct {
	ID uint `param:"id" validate:"required"`
}

type DeleteCustomerRequest struct {
	ID uint `param:"id" validate:"required"`
}
