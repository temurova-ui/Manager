package model

import(
	"time"
)

type Task struct {
	ID          uint `gorm:"primeryKey"`
	Title       string `gorm:"size:100"`
	Description string `gorm:"size:100"`
	Status      string `gorm:"size:50"`
	CreatedAt   time.Time 
   }