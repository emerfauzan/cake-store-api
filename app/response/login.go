package response

type LoginResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
