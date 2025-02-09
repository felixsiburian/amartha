package domain

import "time"

type Loan struct {
	ID            uint `gorm:"primaryKey"`
	CustomerID    uint `gorm:"not null"`
	TotalAmount   int  `gorm:"not null"`
	WeeklyPayment int  `gorm:"not null"`
	Outstanding   int  `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Payments      []Payment `gorm:"foreignKey:LoanID"`
}

type Payment struct {
	ID     uint      `gorm:"primaryKey"`
	LoanID uint      `gorm:"not null"`
	PaidAt time.Time `gorm:"not null"`
	Amount int       `gorm:"not null"`
}
