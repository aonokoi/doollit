package dto

type CreatTaskOutput struct {
	ID int `json:"id"`
}

type CreateTaskInput struct {
	Desc        string `json:"desc"`
	CreatorName string `json:"creator_name"`
}
