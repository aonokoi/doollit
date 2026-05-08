package dto

type CreatTaskOutput struct {
	ID int `json:"id"`
}

type CreateTaskInput struct {
	Desc        string `json:"desc"`
	CreatedName string `json:"created_name"`
}
