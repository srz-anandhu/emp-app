package domain

type Admin struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Role     string `gorm:"column:role;default:'admin'"`
}
