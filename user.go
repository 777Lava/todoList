package todo

type User struct {
	Id       int    `json:"-"` // при сериализации структуры в формат JSON это поле будет игнорироваться.
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
