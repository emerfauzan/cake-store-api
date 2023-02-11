package response

type CakeResponse struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	ImageUrl    string  `json:"image_url"`
}
