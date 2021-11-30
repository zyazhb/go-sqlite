# Pure-Go GORM Sqlite driver
Pure-go (without cgo) implementation of SQLite driver for [GORM](https://gorm.io)

## Usage

```go
import (
  "github.com/glebarez/sqlite"
  "gorm.io/gorm"
)

db, err := gorm.Open(sqlite.Open("file:sqlite.db"), &gorm.Config{})
```

### In-memory DB example
```go
db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
```

### Foreign-key constraint activation
Foreign-key constraint is disabled by default in SQLite. To activate it, use connection parameter:
```go
db, err := gorm.Open(sqlite.Open("file::memory:?_pragma=foreign_keys(1)"), &gorm.Config{})
```
More info: [https://www.sqlite.org/foreignkeys.html](https://www.sqlite.org/foreignkeys.html)

### Shared cache
You also might want to enable shared cache:
```go
db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
```
More info: [https://www.sqlite.org/sharedcache.html](https://www.sqlite.org/sharedcache.html)

