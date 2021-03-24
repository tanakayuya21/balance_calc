package Entity

import (
  "github.com/jinzhu/gorm"
)
type UserBalance struct {
    gorm.Model
    Name   string
    Balance int
}