package main

import (
	"github.com/segment-sources/mysql"
	"github.com/segment-sources/sqlsource"
	"github.com/segment-sources/sqlsource/driver"
)

func main() {
	sqlsource.Run(&driver.MySQL{})
}
