package handler

import (
	"amartha/internal/usecase"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type LoanHandler struct {
	loanUsecase usecase.LoanUsecase
}

func NewLoanHandler(e *echo.Echo, uc usecase.LoanUsecase) {
	handler := &LoanHandler{uc}

	e.POST("/loan", handler.CreateLoan)
	e.GET("/loan/:id/outstanding", handler.GetOutstanding)
	e.POST("/loan/:id/payment", handler.MakePayment)
	e.GET("/loan/:id/delinquent", handler.IsDelinquent)
	e.GET("/loan/:id", handler.GetLoan)
}

func (h *LoanHandler) CreateLoan(c echo.Context) error {
	request := struct {
		CustomerID  uint `json:"customer_id"`
		TotalAmount int  `json:"total_amount"`
	}{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	loan, err := h.loanUsecase.CreateLoan(request.CustomerID, request.TotalAmount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, loan)
}

func (h *LoanHandler) GetOutstanding(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	outstanding, err := h.loanUsecase.GetOutstanding(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int{"outstanding": outstanding})
}

func (h *LoanHandler) MakePayment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	amount := struct {
		Amount int `json:"amount"`
	}{}
	if err := c.Bind(&amount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	err := h.loanUsecase.MakePayment(uint(id), amount.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "payment successful"})
}

func (h *LoanHandler) IsDelinquent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delinquent, err := h.loanUsecase.IsDelinquent(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]bool{"is_delinquent": delinquent})
}

func (h *LoanHandler) GetLoan(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	loan, err := h.loanUsecase.GetLoan(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, loan)
}
