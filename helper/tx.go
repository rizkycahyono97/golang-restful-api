package helper

import (
	"database/sql"
)

// Func untuk helper commit or rollback
func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollBack := tx.Rollback()
		PanicIfError(errorRollBack)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
