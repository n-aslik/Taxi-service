package models

import "time"

type Order struct {
	ID            int       `gorm:"primary key" json:"-"`
	ClientPhone   string    `gorm:"not null" json:"client_phone"`
	ClientAddress string    `gorm:"not null" json:"client_address"`
	From          string    `gorm:"not null" json:"from"`
	Into          string    `gorm:"not null" json:"into"`
	Distance      int       `json:"distance"`
	StartPrice    int       `json:"start_price"`
	AllPrice      int       `json:"all_price"`
	DriverPhone   string    `gorm:"not null" json:"driver_phone"`
	ClientID      int       `gorm:"references users(id)" json:"-"`
	DriverID      int       `gorm:"references users(id)" json:"-"`
	IsResponse    bool      `gorm:"default false" json:"is_response"`
	IsDeleted     bool      `gorm:"default false" json:"-"`
	CreatedAt     time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"-" gorm:"autoUpdateTime"`
}

func (Order) TableName() string {
	return "orders"
}

type SendRoute struct {
	From string `gorm:"not null" json:"from"`
	Into string `gorm:"not null" json:"into"`
}

type EditOrder struct {
	DriverPhone string `gorm:"not null" json:"driver_phone"`
	DriverID    int    `gorm:"references users(id)" json:"driver_id"`
}
type AddDistances struct {
	Distance int `gorm:"not null" json:"distance"`
}

type GetOrder struct {
	ClientPhone string `gorm:"not null" json:"client_phone"`
	From        string `gorm:"not null" json:"from"`
	Into        string `gorm:"not null" json:"into"`
	Distance    int    `gorm:"not null" json:"distance"`
	StartPrice  int    `json:"start_price"`
	AllPrice    int    `json:"all_price"`
	DriverPhone string `gorm:"not null" json:"driver_phone"`
	ClientID    int    `gorm:"references users(id)" json:"-"`
	DriverID    int    `gorm:"references users(id)" json:"-"`
	IsResponse  bool   `gorm:"default false" json:"is_response"`
	IsDeleted   bool   `gorm:"default false" json:"-"`
}

type Reports struct {
	From       string    `gorm:"not null" json:"from"`
	Into       string    `gorm:"not null" json:"into"`
	Distance   int       `gorm:"not null" json:"distance"`
	StartPrice int       `json:"start_price"`
	AllPrice   int       `json:"all_price"`
	ClientID   int       `gorm:"references users(id)" json:"client_id"`
	DriverID   int       `gorm:"references users(id)" json:"driver_id"`
	IsResponse bool      `gorm:"default false" json:"is_response"`
	IsDeleted  bool      `gorm:"default false" json:"is_deleted"`
	CreatedAt  time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"-" gorm:"autoUpdateTime"`
}

type Checkresponse struct {
	IsResponse bool `json:"is_response"`
}
