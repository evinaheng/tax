package repository

import (
	"log"

	"github.com/tax/internal/entity"
	"github.com/tax/lib/storage/database"
)

// NewBills repository for system
func NewBills(db database.Database) Bills {
	return &billsRepo{
		db: db,
	}
}

func (r *billsRepo) CreateBills(param entity.Bills) error {

	tx, err := r.db.Begin()
	if err != nil {
		log.Println("func CreateBills", err)
		return err
	}

	query := `
		INSERT INTO bills(name,tax_code,price)
		VALUES(?,?,?)
	`

	_, err = tx.Exec(query, param.Name, param.TaxCode, param.Price)
	if err != nil {
		tx.Rollback()
		log.Println("func CreateBills", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		log.Println("func CreateBills", err)
		return err
	}

	return err
}

func (r *billsRepo) GetBills() []entity.Bills {

	res := []entity.Bills{}

	query := `
	SELECT
		id, 
		name, 
		tax_code, 
		price
	FROM bills
	`

	// Query all bills list from db
	rows, err := r.db.Queryx(query)
	if err != nil {
		log.Println("func GetBills", err)
		return res
	}

	defer rows.Close()
	for rows.Next() {
		eb := entity.Bills{}
		rows.StructScan(&eb)

		res = append(res, eb)
	}
	return res

}
