package models

type Port struct {
	Port       string `json:"port" gorm:"unique"`
	ClientName string `json:"clientname" gorm:"unique"`
}
