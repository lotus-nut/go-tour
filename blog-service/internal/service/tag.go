package service

type CountTagRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=1" verbose_name:"名称"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1" verbose_name:"状态值"`
}

type CreateTagRequest struct {
	Name     string `form:"name" binding:"max=100,min=3,required"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreateBy string `form:"created_by" binding:"max=100,min=3,required"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"gte=1,required"`
	Name       string `form:"name" binding:"max=100,min=3"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1,required"`
	ModifiedBy string `form:"modified_by" binding:"max=100,min=3,required"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"gte=1,required"`
}
