package usersdto

type CreateUserReq struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string `json:"role" form:"role" validate:"required"`
}

type UpdateUserReq struct {
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

type UserRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
