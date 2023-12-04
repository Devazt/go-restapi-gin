package auth

type AuthReq struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginReq struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginRes struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
