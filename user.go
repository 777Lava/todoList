package todo

type User struct {
	Id       int    `json:"-" db:"id"` // при сериализации структуры в формат JSON это поле будет игнорироваться.
	Name     string `json:"name" binding:"required"` 
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
