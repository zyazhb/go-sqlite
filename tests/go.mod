module github.com/glebarez/sqlite/tests

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/glebarez/sqlite v1.2.9
	github.com/google/uuid v1.3.0
	github.com/jinzhu/now v1.1.3
	github.com/lib/pq v1.10.4
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.0.0-20211204120058-94396e421777 // indirect
	golang.org/x/tools v0.1.8 // indirect
	gorm.io/gorm v1.22.5-0.20211202023924-300a23fc3137
	modernc.org/ccgo/v3 v3.12.86 // indirect
	modernc.org/sqlite v1.14.3-0.20211203211519-cbb9557100f0 // indirect
)

replace github.com/glebarez/sqlite => ../
