module github.com/glebarez/sqlite/tests

go 1.16

require (
	github.com/glebarez/sqlite v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
	github.com/jinzhu/now v1.1.4
	github.com/lib/pq v1.10.4
	gorm.io/gorm v1.22.5-0.20211202023924-300a23fc3137
)

replace github.com/glebarez/sqlite => ../
