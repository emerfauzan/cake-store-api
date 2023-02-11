package request

type CreateCakeRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	ImageUrl    string  `json:"image_url"`
	UserId      uint    `json:"-"`
}

type UpdateCakeRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	ImageUrl    string  `json:"image_url"`
	UserId      uint    `json:"-"`
}

type GetCakesRequest struct {
	Offset  uint
	Page    uint
	Limit   uint
	Sorts   []string
	Roles   []string
	Keyword string
}
