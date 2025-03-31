package models

import ()

// List Model
type List struct {
    ListKey  string `gorm:"primaryKey"`
    Value    string
    Position int `gorm:"primaryKey"`
}

// Set Model
type Set struct {
    SetKey string `gorm:"primaryKey"`
    Value  string `gorm:"primaryKey"`
}

// Hash Model
type Hash struct {
    HashKey string `gorm:"primaryKey"`
    Field   string `gorm:"primaryKey"`
    Value   string
}
