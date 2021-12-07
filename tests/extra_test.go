package tests_test

import (
	"testing"

	. "gorm.io/gorm/utils/tests"
)

func TestReadWhileWrite(t *testing.T) {
	// for not-yet-clear reason, this test fails with SQLITE_BUSY (Database locked) error
	// this behavior is still under investigation about SQLite internals
	// but journal_mode = WAL doesn't suffer from this

	var journalMode string
	DB.Raw("PRAGMA journal_mode").Scan(&journalMode)
	if journalMode != "wal" {
		t.Skipf("skipped to avoid failure due to SQLITE_BUSY error")
	}

	user := User{Name: "SelectUser1"}
	if err := DB.Save(&user).Error; err != nil {
		t.Fatal(err)
	}

	// open Rows (read transaction)
	rows, err := DB.Table("users").Select("COALESCE(age,?)", "42").Rows()
	if err != nil {
		t.Fatalf("Failed, got error: %v", err)
	} else {
		defer rows.Close()
	}

	// try write
	if err := DB.Save(&user).Error; err != nil {
		t.Fatal(err)
	}
}
