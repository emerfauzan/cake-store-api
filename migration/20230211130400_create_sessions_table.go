package migration

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20230211130400",
		Up:      mig_20230211130400_create_sessions_table_up,
		Down:    mig_20230211130400_create_sessions_table_down,
	})
}

func mig_20230211130400_create_sessions_table_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE sessions (id serial primary key, user_id int, session_key varchar(255), expires_at datetime);")
	if err != nil {
		return err
	}
	return nil
}

func mig_20230211130400_create_sessions_table_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE sessions")
	if err != nil {
		return err
	}
	return nil
}
