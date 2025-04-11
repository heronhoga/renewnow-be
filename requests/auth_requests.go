package requests

type CreateUserRequest struct {
    Username  string    `json:"username" validate:"required,min=1"`
    Password  string    `json:"password" validate:"required,min=8"`
    FirstName string    `json:"first_name" validate:"required,min=1"`
    LastName  string    `json:"last_name"`
    Phone     string    `json:"phone" validate:"required,number,startswith=62"`
}

type LoginUserRequest struct {
    Username string `json:"username" validate:"required,min=1"`
    Password string `json:"password" validate:"required,min=1"`
}

type LogoutUserRequest struct {
    Username string `json:"username" validate:"required,min=1"`
}
