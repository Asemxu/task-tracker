package models

import "fmt"

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

func (c Task) ShowTask() {
	fmt.Printf("Id:  %v\n", c.Id)
	fmt.Printf("Description:  %v\n", c.Description)
	fmt.Printf("Status:  %v\n", c.Status)
	fmt.Printf("CreatedAt:  %v\n", c.CreatedAt)
	fmt.Printf("UpdatedAt:  %v\n", c.UpdatedAt)

}
