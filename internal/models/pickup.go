package models

type Pickup struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	PickupTime string `json:"time"`
}
