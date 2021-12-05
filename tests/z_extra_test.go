package tests_test

import (
	"testing"
)

func TestReadWhileWrite(t *testing.T) {
	// t.Skip()
	user := User{Name: "SelectUser1"}
	if err := DB.Save(&user).Error; err != nil {
		t.Fatal(err)
	}

	// open Rows (read transaction)
	rows, err := DB.Table("users").Select("COALESCE(age,?)", "42").Rows()
	if err != nil {
		t.Fatalf("Failed, got error: %v", err)
	}
	defer rows.Close()

	// try write
	if err := DB.Save(&user).Error; err != nil {
		t.Fatal(err)
	}
}
