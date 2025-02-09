package repository

import (
	"amartha/internal/domain"
	"github.com/jinzhu/gorm"
)

type LoanRepository interface {
	CreateLoan(loan *domain.Loan) error
	GetLoanByID(id uint) (*domain.Loan, error)
	UpdateLoan(loan *domain.Loan) error
	GetPaymentsByLoanID(loanID uint) ([]domain.Payment, error)
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db}
}

func (r *loanRepository) CreateLoan(loan *domain.Loan) error {
	return r.db.Create(loan).Error
}

func (r *loanRepository) GetLoanByID(id uint) (*domain.Loan, error) {
	var loan domain.Loan
	err := r.db.Preload("Payments").First(&loan, id).Error
	return &loan, err
}

func (r *loanRepository) UpdateLoan(loan *domain.Loan) error {
	return r.db.Save(loan).Error
}

func (r *loanRepository) GetPaymentsByLoanID(loanID uint) ([]domain.Payment, error) {
	var payments []domain.Payment
	if err := r.db.Where("loan_id = ?", loanID).Order("paid_at ASC").Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
