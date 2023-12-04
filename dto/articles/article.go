package articlesdto

type CreateArticleReq struct {
	Title   string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
	Image   string `json:"image" form:"image"`
	UserID  int    `json:"user_id" form:"user_id"`
}

type UpdateArticleReq struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Image   string `json:"image" form:"image"`
	UserID  int    `json:"user_id" form:"user_id"`
}

type ArticleRes struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}
