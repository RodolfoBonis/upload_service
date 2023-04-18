package models

type IpModel struct {
	Address string `json:"address,omitempty" validate:"required"`
	Health  bool   `json:"health,omitempty" validate:"required"`
}
