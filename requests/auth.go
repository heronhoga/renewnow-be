package requests

type CreateUserRequest struct {
    Username  string    `json:"username" validate:"required,min=1"`
    Password  string    `json:"password" validate:"required,min=8"`
    FirstName string    `json:"first_name" validate:"required,min=1"`
    LastName  string    `json:"last_name"`
    Email     string    `json:"email" validate:"required,email"`
}

type LoginUserRequest struct {
    Username string `json:"username" validate:"required,min=1"`
    Password string `json:"password" validate:"required,min=1"`
}

type LogoutUserRequest struct {
    Username string `json:"username" validate:"required,min=1"`
}
