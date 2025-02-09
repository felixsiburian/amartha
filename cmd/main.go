package main

import (
	"amartha/config"
	"amartha/internal/handler"
	"amartha/internal/repository"
	"amartha/internal/usecase"
	"github.com/labstack/echo"
)

func main() {
	app := config.Config{}

	app.CatchError(app.InitEnv())

	dbConfig := app.GetDBConfig()

	db := config.ConnectionDB(dbConfig)

	loanRepo := repository.NewLoanRepository(db)
	loanUsecase := usecase.NewLoanUsecase(loanRepo)

	e := echo.New()
	handler.NewLoanHandler(e, loanUsecase)

	e.Start(":8080")
}
