package sqlite

import "fmt"

// defaultOptions are set when calling Open method, may be redefined with Open() options parameter
var defaultOptions = []Option{
	ForeignKeysOn,
	JournalModeWAL,
	BusyTimeout(5000),
}

// some predefined values
var (
	ForeignKeysOn  = ForeignKeys(true)
	ForeignKeysOff = ForeignKeys(false)
	JournalModeWAL = JournalMode("WAL")
)

// https://sqlite.org/pragma.html#pragma_foreign_keys
func ForeignKeys(on bool) Option {
	return BoolPragma(`foreign_keys`, on)
}

// https://sqlite.org/pragma.html#pragma_journal_mode
func JournalMode(mode string) Option {
	return &Pragma{
		Name:  `journal_mode`,
		Value: mode,
	}
}

// https://sqlite.org/c3ref/busy_timeout.html
func BusyTimeout(millis int) Option {
	return &Pragma{
		Name:  `busy_timeout`,
		Value: millis,
	}
}

type Option interface{}

type Pragma struct {
	Option
	Name  string
	Value interface{}
}

func BoolPragma(name string, value bool) Option {
	return &Pragma{
		Name:  name,
		Value: fmt.Sprint(value),
	}
}
