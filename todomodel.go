package  main
import (
	"time"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type (
	// todoModel describes a todoModel type
	TodoModel struct {
	ID   int `gorm:"AUTO_INCREMENT"`
	CreatedAt time.Time
    Title     string `json:"title"`
	 Completed int    `json:"completed"`
	 Comment string `json: "comment"`
	}
	Comment struct {
		ID string
		Comment string
	}
   )