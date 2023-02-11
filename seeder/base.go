package seeder

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type Seeder struct {
	db *sql.DB
}

type SeederData struct {
	Table string
	Data  []map[string]interface{}
}

func Init(db *sql.DB) (*Seeder, error) {
	var seeder = &Seeder{
		db: db,
	}

	return seeder, nil
}

func (s *Seeder) Seed() error {
	f, err := os.Open("./seeder/data.json")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	var seeders []SeederData

	json.Unmarshal(b, &seeders)
	tx, err := s.db.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		fmt.Println(err)
	}

	for _, seeder := range seeders {
		for _, data := range seeder.Data {
			columnName := ""
			valueName := ""
			count := 0

			for key, value := range data {
				if key == "encrypted_password" {
					value, _ = bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%v", value)), 14)
				}

				if count == 0 {
					columnName += fmt.Sprintf("(%s,", key)
					valueName += fmt.Sprintf("('%s',", value)
				} else if count == len(data)-1 {
					columnName += fmt.Sprintf("%s)", key)
					valueName += fmt.Sprintf("'%s')", value)
				} else {
					columnName += fmt.Sprintf("%s,", key)
					valueName += fmt.Sprintf("'%s',", value)
				}

				count++

			}

			query := fmt.Sprintf("INSERT INTO %s %s values %s", seeder.Table, columnName, valueName)

			if _, err := tx.Exec(query); err != nil {
				tx.Rollback()
				return err
			}

		}
	}

	tx.Commit()

	fmt.Println("Database seeders are seeded succesfully")

	return nil

}
