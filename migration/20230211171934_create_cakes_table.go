package migration

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20230211171934",
		Up:      mig_20230211171934_create_cakes_table_up,
		Down:    mig_20230211171934_create_cakes_table_down,
	})
}

func mig_20230211171934_create_cakes_table_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE cakes (id serial primary key, title varchar(100), description varchar(500), rating double, image_url varchar(255), is_deleted_flag boolean, created_by int, updated_by int, deleted_by int, created_at datetime, updated_at datetime, deleted_at datetime);")
	if err != nil {
		return err
	}
	return nil
}

func mig_20230211171934_create_cakes_table_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE cakes")
	if err != nil {
		return err
	}
	return nil
}
