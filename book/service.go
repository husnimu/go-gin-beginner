package book

type BookRequest struct {
	Title       string `binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Description string `binding:"required"`
	Rating      int    `binding:"required,number"`
}
