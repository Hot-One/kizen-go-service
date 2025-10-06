package dto

import "github.com/Hot-One/kizen-go-service/pkg/pg"

type UserPage = pg.PageData[User] // @name UserPage

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at" swaggerignore:"true"`
	UpdatedAt string `json:"updated_at" swaggerignore:"true"`
} // @name User

type CreateUser struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
} // @name CreateUser

type UpdateUser struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email" binding:"omitempty,email"`
} // @name UpdateUser
