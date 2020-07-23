package service

type CountArticleRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleRequest struct {
	ID uint32 `form:"id" binding:"gte=1,required"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"max=100,min=3,required"`
	Desc          string `form:"desc" binding:"max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=255"`
	Content       string `form:"content" binding:"required"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreateBy      string `form:"created_by" binding:"max=100,min=3,required"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"gte=1,required"`
	Title         string `form:"title" binding:"max=100,min=3"`
	Desc          string `form:"desc" binding:"max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=255"`
	Content       string `form:"content" binding:"required"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1,required"`
	ModifiedBy    string `form:"modified_by" binding:"max=100,min=3,required"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"gte=1,required"`
}
