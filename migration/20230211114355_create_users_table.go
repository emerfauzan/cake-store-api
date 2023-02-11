package migration

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20230211114355",
		Up:      mig_20230211114355_create_users_table_up,
		Down:    mig_20230211114355_create_users_table_down,
	})
}

func mig_20230211114355_create_users_table_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE users ( name varchar(255) );")
	if err != nil {
		return err
	}
	return nil
}

func mig_20230211114355_create_users_table_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users")
	if err != nil {
		return err
	}
	return nil
}
