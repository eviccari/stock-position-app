package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/eviccari/stock-position-app/adapters/repositories"
	"github.com/eviccari/stock-position-app/app/services"
	"github.com/eviccari/stock-position-app/app/vos"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	db, err := sql.Open("mysql", "stock_pos_user_service:123@tcp(localhost:3306)/products?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	repo := repositories.NewMySQLRepository(db)
	service := services.NewStockService(repo)

	sp, err := service.Update(vos.StockMovVO{
		ID:                     uuid.NewString(),
		ItemID:                 "item0101010",
		FacilityID:             30,
		FromAvailQty:           0,
		ToAvailQty:             10,
		FromUnavlQty:           0,
		ToUnavlQty:             0,
		TransactionDescription: "unavl for cycle count",
		TransactionID:          uuid.NewString(),
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(sp)

	sp, err = service.Update(vos.StockMovVO{
		ID:                     uuid.NewString(),
		ItemID:                 "item0101010",
		FacilityID:             30,
		FromAvailQty:           10,
		ToAvailQty:             0,
		FromUnavlQty:           0,
		ToUnavlQty:             10,
		TransactionDescription: "unavl for cycle count",
		TransactionID:          uuid.NewString(),
	})

	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(sp)
	os.Exit(0)
}
