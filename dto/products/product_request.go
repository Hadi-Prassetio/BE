package productsdto

type CreateProductRequest struct {
	Title string `json:"title" gorm:"type: varchar(255)" validate:"required"`
	Price int    `json:"price" gorm:"type: varchar(255)" validate:"required"`
}

type UpdateProductRequest struct {
	Title string `json:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" gorm:"type: varchar(255)"`
}
