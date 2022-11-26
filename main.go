package main

import (
	"fmt"

	"github.com/abhishek2966/singleton-demo/pkg/ops"
	"github.com/abhishek2966/singleton-demo/pkg/pgdb"
)

func main() {
	counter := ops.GetCounterInstance()
	counter.AddOne()

	fmt.Println(counter.GetCount())

	//new connection is opened
	db, err := pgdb.GetDBInstance()
	// do some query work on connection
	fmt.Println(db, err)
	// connection already opened is returned
	db, err = pgdb.GetDBInstance()
	fmt.Println(db, err)
}
