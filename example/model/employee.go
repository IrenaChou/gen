package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type Employee struct {
	EmpNo     int         `gorm:"column:emp_no;primary_key" json:"emp_no"`
	BirthDate null.Time   `gorm:"column:birth_date" json:"birth_date"`
	FirstName null.String `gorm:"column:first_name" json:"first_name"`
	LastName  null.String `gorm:"column:last_name" json:"last_name"`
	Gender    null.String `gorm:"column:gender" json:"gender"`
	HireDate  null.Time   `gorm:"column:hire_date" json:"hire_date"`
}

// TableName sets the insert table name for this struct type
func (e *Employee) TableName() string {
	return "employees"
}
