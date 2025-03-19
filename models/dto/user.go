package dto

type RegisterMap struct {
	UserName string `json:"userName" gorm:"column:userName"`
	Password string `json:"password,omitempty" gorm:"column:password"` //omitempty表示可选
	Email    string `json:"email" gorm:"column:email"`
	Avatar   string `json:"avatar" gorm:"avatar"`
}
