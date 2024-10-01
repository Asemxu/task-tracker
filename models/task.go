package models

type Task struct {
	Id          float64 `json:"id"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

func (c Task) ID() (jsonField string, value interface{}) {
	value = c.Id
	jsonField = "id"
	return
}
