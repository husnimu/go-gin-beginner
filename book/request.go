package book

type UpdateBookRequest struct {
	Title       string
	Price       int
	Description string
	Rating      int
}

type BookRequest struct {
	Title       string `binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Description string `binding:"required"`
	Rating      int    `binding:"required,number"`
}
