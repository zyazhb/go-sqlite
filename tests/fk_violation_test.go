package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const dsn = "file::memory:?cache=shared&_pragma=foreign_keys(1)"

var (
	// this query must produce database error due to foreign key constraint violation
	violationQuery = "INSERT INTO `child` (`parent_id`) VALUES (\"non-existing\") RETURNING `id`"

	// gorm config
	config = &gorm.Config{
		// for debugging you may set logging level
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,

		// singular table name
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
)

type Parent struct {
	ID string `gorm:"primaryKey"`
}

type Child struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement;not null"`
	ParentID string
	Parent   Parent
}

var (
	db           *gorm.DB
	validChildID uint64
)

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(sqlite.Open(dsn), config)
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}

	//migrate
	if err := db.Migrator().DropTable(&Parent{}, &Child{}); err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&Parent{}, &Child{}); err != nil {
		log.Fatal(err)
	}

	// create valid records
	child := &Child{
		Parent: Parent{ID: "valid-parent"},
	}
	if err := db.Create(child).Error; err != nil {
		log.Fatal(err)
	}
	validChildID = child.ID
	fmt.Printf("valid child ID: %d\n", validChildID)

	// run tests
	os.Exit(m.Run())
}

func Test_Create(t *testing.T) {
	require := require.New(t)

	// create child for non-existing parent
	child := &Child{
		ParentID: "non-existing",
	}

	err := db.Create(child).Error
	require.Error(err)
	require.Equal(err.Error(), "constraint failed: FOREIGN KEY constraint failed (787)")
}

func Test_Exec(t *testing.T) {
	require := require.New(t)
	err := db.Exec(violationQuery).Error
	require.Error(err)
	require.Equal(err.Error(), "constraint failed: FOREIGN KEY constraint failed (787)")
}

func Test_Update(t *testing.T) {
	require := require.New(t)

	// create child for non-existing parent
	err := db.Updates(&Child{
		ID:       validChildID,
		ParentID: "non-existing",
	}).Error
	require.Error(err)
	require.Equal(err.Error(), "constraint failed: FOREIGN KEY constraint failed (787)")
}
