package migration

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20230211123347",
		Up:      mig_20230211123347_create_cakes_table_up,
		Down:    mig_20230211123347_create_cakes_table_down,
	})
}

func mig_20230211123347_create_cakes_table_up(tx *sql.Tx) error {
	return nil
}

func mig_20230211123347_create_cakes_table_down(tx *sql.Tx) error {
	return nil
}