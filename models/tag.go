package models

type Tag struct {
	Model
	Name       string `json:"name" binding:"required"`
	CreatedBy  string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state" binding:"required,eq=1|eq=2"`
}

type QueryTag struct {
	Name  string `form:"name"`
	State int    `form:"state" binding:"omitempty,eq=1|eq=2"`
}

/*func (q QueryTag) TableName() string {
	return "tag" // find方法会添加前缀，where却不会
}*/
