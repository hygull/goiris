/**
@file name 	 : callees.go

@description : This is the file containing the definitions of functions used by other
			   functions of controllers package
*/

package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"time"
)
import (
	"goiris/conf"
	"goiris/views"
)

func SetContentTypeAndShowPath(ctx *iris.Context) {
	ctx.SetContentType("text/html")
	fmt.Println("Serving : ", ctx.PathString())
}

func OpenDbConn() (*sql.DB, error) {
	connString := conf.DbUser + ":" + conf.DbPassword + "@tcp(" + conf.DbHost + ":" + conf.DbPort + ")/" + conf.Db
	fmt.Println("\n[SynkkU] Opening a DB connection with ", connString)
	db, err := sql.Open("mysql", connString)
	return db, err
}

func CheckDbConn(ctx *iris.Context, err error) bool {
	t := time.Now().String()[0:19]
	if err != nil {
		fmt.Println("\n[SynkkU] "+t+" Can't open a DB connection. Check DB credentials then retry...", err)
		views.ShowErrorAsJSON(ctx, 6)
		return false
	} else {
		fmt.Println("\n[SynkkU] " + t + " DB connection successfully opened...")
		return true
	}
}

func CheckForEmptyFields(ctx *iris.Context, urlKeys ...string) (isEmpty bool) {
	for _, key := range urlKeys {
		if key == "" {
			views.ShowErrorAsJSON(ctx, 7)
			isEmpty = true
		}
	}
	if isEmpty {
		fmt.Println("\n[SynkkU] Either any field/key don't have its value or you have provided wrong key names...")
	} else {
		fmt.Println("\n[SynkkU] All field's/key's names are OK and they have values...")
	}
	return //It will return the value contained in isEmpty
}

var queryTypes = []string{"insert", "select", "update", "delete"}

func ExecuteQuery(ctx *iris.Context, db *sql.DB, query string, index int) (*sql.Rows, int64) {
	var err error
	var stmt *sql.Stmt
	var rows *sql.Rows

	if index == 2 {
		rows, err = db.Query(query)
	} else {
		if (index <= 4) && (index >= 1) {
			stmt, err = db.Prepare(query)
		} else {
			views.ShowErrorAsJSON(ctx, 8)
			return nil, 0
		}
	}

	var id int64
	if err != nil {
		views.ShowErrorAsJSON(ctx, index)
		return nil, id
	} else {
		if index != 2 {
			res, e := stmt.Exec()
			if e != nil {
				fmt.Println("\n[SynkkU] Error while executing Exec()")
				return nil, 0
			}
			if index == 1 {
				id, e = res.LastInsertId()
				if e != nil {
					fmt.Println("\n[SynkkU] Error while executing Exec()")
					return nil, 0
				}
			}
		}
		fmt.Println("\n[SynkkU] " + queryTypes[index-1] + " query successfully executed...")
	}
	if index == 2 {
		return rows, id
	}
	return nil, id
}

func CheckForUserExistance(ctx *iris.Context, db *sql.DB, uid string) (exists bool, status int) {
	status = 0
	rows, err := db.Query("select user_id from users where user_id=" + uid + ";")
	if err != nil {
		fmt.Println("\n[SynkkU] Error in execution of select query...")
		views.ShowErrorAsJSON(ctx, 2)
		exists = false
		status = 1
		return
	}
	if rows.Next() {
		exists = true
		return
	}
	exists = false
	return
}
