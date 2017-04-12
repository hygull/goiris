package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"strconv"
)

// import (
// 	"goiris/views"
// )

func PostOffice(ctx *iris.Context) {
	fmt.Println("Got a request for a PostOffice")
	pin, err := ctx.ParamInt("pincode")

	println(pin, err, err == nil, err != nil)
	if err != nil {
		ctx.JSON(iris.StatusBadRequest, iris.Map{"status": 400, "msg": "Invalid pincode"})
		return
	}

	SetContentTypeAndShowPath(ctx)
	db, err := OpenDbConn()
	isDbConnOK := CheckDbConn(ctx, err)
	if !isDbConnOK {
		return
	}

	defer db.Close()

	// query := "insert into feedback set user_id=" + uid + ", fback_message='" + feedback_message + "'"
	query := "select pincode,city, state, district from PyApp_Postoffice where id=" + strconv.Itoa(pin) + ";"
	fmt.Println("\n[SynkkU] Executing : ", query)

	rows, id := ExecuteQuery(ctx, db, query, 2) //2nd parameter denotes the status_code-100...eg. for 101 we have to give 1
	if rows == nil && id == 0 {
		return
	}
	if rows != nil {
		fmt.Println("\n[SynkkU] feedback successfully posted...")
	}

	if id != 0 {
		fmt.Println("\n[SynkkU] Last inserted feedback id : ", id)
	}

	// views.ShowJSON(ctx, []string{"success", "message", "febk_id"}, []interface{}{1, "Your feedback sent successfully", id})
	type PostOfficeData struct {
		Pincode  int    `json:"pincode"`
		City     string `json:"city"`
		State    string `json:"state"`
		District string `json:"district"`
	}

	PostOffice := PostOfficeData{}

	var pincode int
	var state string
	var city string
	var district string

	found := false
	for rows.Next() {
		found = true
		rows.Scan(&pincode, &city, &state, &district)

		PostOffice = PostOfficeData{pincode, city, state, district}
	}

	if found {
		ctx.JSON(200, PostOffice) //return is not required as it is the last statement
	} else {
		ctx.JSON(404, iris.Map{"status": 404, "msg": "Data not found"})
	}

}
