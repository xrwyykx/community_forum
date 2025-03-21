package dto

type RegisterMap struct {
	UserName string `json:"userName" gorm:"column:userName"`
	Password string `json:"password,omitempty" gorm:"column:password"` //omitempty表示可选
	Email    string `json:"email" gorm:"column:email"`
	Avatar   string `json:"avatar" gorm:"avatar"`
}

type LoginMap struct {
	UserName string `json:"userName" gorm:"column:userName"`
	Password string `json:"password,omitempty" gorm:"column:password"` //omitempty表示可选
}

//type InfoMap struct {
//	UserId int64 `json:"userId"`
//}

type GetInfoMap struct {
	UserName string `json:"userName" gorm:"column:userName"`
	Email    string `json:"email" gorm:"column:email"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
	Points   int    `json:"point" gorm:"column:points"`
}

type UpdateInfoMap struct {
	UserId   int64  `json:"userId" gorm:"column:userId"`
	UserName string `json:"userName,omitempty" gorm:"column:userName"`
	Password string `json:"password,omitempty" gorm:"column:password"`
	Email    string `json:"email,omitempty" gorm:"column:email"`
	Avatar   string `json:"avatar,omitempty" gorm:"column:avatar"`
}
