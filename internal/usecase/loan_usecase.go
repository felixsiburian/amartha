package usecase

import (
	"amartha/internal/domain"
	"amartha/internal/repository"
	"errors"
	"time"
)

type LoanUsecase interface {
	CreateLoan(customerID uint, totalAmount int) (*domain.Loan, error)
	GetOutstanding(id uint) (int, error)
	MakePayment(id uint, amount int) error
	IsDelinquent(id uint) (bool, error)
	GetLoan(id uint) (*domain.Loan, error)
}

type loanUsecase struct {
	loanRepo repository.LoanRepository
}

func NewLoanUsecase(loanRepo repository.LoanRepository) LoanUsecase {
	return &loanUsecase{loanRepo}
}

func (uc *loanUsecase) CreateLoan(customerID uint, totalAmount int) (*domain.Loan, error) {
	weeks := 50
	interestRate := 0.1
	totalWithInterest := int(float64(totalAmount) * (1 + interestRate))
	weeklyPayment := totalWithInterest / weeks

	loan := &domain.Loan{
		CustomerID:    customerID,
		TotalAmount:   totalWithInterest,
		WeeklyPayment: weeklyPayment,
		Outstanding:   totalWithInterest,
	}

	err := uc.loanRepo.CreateLoan(loan)
	if err != nil {
		return nil, err
	}
	return loan, nil
}

func (uc *loanUsecase) GetOutstanding(id uint) (int, error) {
	loan, err := uc.loanRepo.GetLoanByID(id)
	if err != nil {
		return 0, err
	}
	return loan.Outstanding, nil
}

func (uc *loanUsecase) MakePayment(id uint, amount int) error {
	loan, err := uc.loanRepo.GetLoanByID(id)
	if err != nil {
		return err
	}

	payments, err := uc.loanRepo.GetPaymentsByLoanID(loan.ID)
	if err != nil {
		return err
	}

	weeksSinceLoanStart := int(time.Since(loan.CreatedAt).Hours()/24/7) + 1
	missedPayments := weeksSinceLoanStart - len(payments)

	if missedPayments >= 2 {
		requiredAmount := missedPayments * loan.WeeklyPayment
		if amount != requiredAmount {
			return errors.New("delinquent borrowers must pay exactly the amount of missed payments")
		}
	}

	loan.Outstanding -= amount
	payment := domain.Payment{LoanID: loan.ID, Amount: amount, PaidAt: time.Now()}
	loan.Payments = append(loan.Payments, payment)

	return uc.loanRepo.UpdateLoan(loan)
}

func (uc *loanUsecase) IsDelinquent(id uint) (bool, error) {
	loan, err := uc.loanRepo.GetLoanByID(id)
	if err != nil {
		return false, err
	}

	missed := 0
	for i := len(loan.Payments) - 1; i >= 0; i-- {
		if time.Since(loan.Payments[i].PaidAt).Hours() > 7*24 {
			missed++
			if missed >= 2 {
				return true, nil
			}
		} else {
			break
		}
	}
	return false, nil
}

func (uc *loanUsecase) GetLoan(id uint) (*domain.Loan, error) {
	return uc.loanRepo.GetLoanByID(id)
}
