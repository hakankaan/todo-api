package store

// Todo is a struct for todo
type Todo struct {
	ID          uint   `gorm:"column:id;primaryKey;autoIncrement" json:"-"`
	GuID        string `gorm:"column:global" json:"global"`
	Title       string `gorm:"column:title;not null" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	IsDone      bool   `gorm:"column:is_done;default:false" json:"is_done"`
}
