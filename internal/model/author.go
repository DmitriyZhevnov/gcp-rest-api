package model

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateAuthorDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateAuthorDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
