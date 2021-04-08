package http_rest_api_test

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
