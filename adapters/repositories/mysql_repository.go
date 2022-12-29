package repositories

import (
	"database/sql"

	"github.com/eviccari/stock-position-app/app/vos"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) Create(vo vos.StockPosVO) (err error) {
	q := `insert into products.stock_positions (
		      id, 
			  item_id, 
			  facility_id, 
			  on_hand_qty, 
			  unavl_qty, 
			  position_date, 
			  transaction_id
			) values (?, ?, ?, ?, ?, ?, ?)`

	_, err = r.db.Exec(q, vo.ID, vo.ItemID, vo.FacilityID, vo.OnHandQty, vo.UnavlQty, vo.PositionDate, vo.TransactionID)
	return
}

func (r *MySQLRepository) Update(vo vos.StockPosVO, constraintKey string) (updatedRowCount int64, err error) {
	q := `update products.stock_positions
		     set on_hand_qty    = ?, 
			     unavl_qty      = ?,
				 position_date  = ?,
				 transaction_id = ?
		   where id             = ?
		     and transaction_id = ?`

	res, err := r.db.Exec(q,
		vo.OnHandQty,
		vo.UnavlQty,
		vo.PositionDate,
		vo.TransactionID,
		vo.ID,
		constraintKey,
	)

	if err != nil {
		return
	}

	updatedRowCount, err = res.RowsAffected()
	return
}

func (r *MySQLRepository) FindByItemIDAndFacilityID(itemID string, facilityID uint) (vo vos.StockPosVO, err error) {
	q := "select * from products.stock_positions where item_id = ? and facility_id = ?"

	row := r.db.QueryRow(q, itemID, facilityID)

	err = row.Scan(&vo.ID, &vo.ItemID, &vo.FacilityID, &vo.OnHandQty, &vo.UnavlQty, &vo.PositionDate, &vo.TransactionID)
	if isEmpty := isEmptyError(err); isEmpty {
		return vos.StockPosVO{}, nil
	}
	return
}

func isEmptyError(err error) bool {
	return err != nil && (err.Error() == sql.ErrNoRows.Error())
}
