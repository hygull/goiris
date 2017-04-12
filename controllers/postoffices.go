package controllers

import (
	"fmt"
	"github.com/kataras/iris"
)

// import (
// 	"goiris/views"
// )

func PostOffices(ctx *iris.Context) {
	SetContentTypeAndShowPath(ctx)
	db, err := OpenDbConn()
	isDbConnOK := CheckDbConn(ctx, err)
	if !isDbConnOK {
		return
	}
	defer db.Close()

	// pincode := strings.TrimSpace(ctx.FormValueString("pincode"))
	// feedback_message := strings.TrimSpace(ctx.FormValueString("febk_mes"))
	// fmt.Println("\n[SynkkU] Gathered infomation : uid="+uid, " and feedback_message="+feedback_message)

	// isEmpty := CheckForEmptyFields(ctx, uid, feedback_message)
	// if isEmpty {
	// 	return
	// }

	// exists, status := CheckForUserExistance(ctx, db, uid)

	// if status == 1 {
	// 	return
	// }

	// if !exists {
	// 	views.ShowErrorAsJSON(ctx, 9)
	// 	fmt.Println("\n[SynkkU] User id is not present on the database, Create an account please...")
	// 	return
	// }

	// query := "insert into feedback set user_id=" + uid + ", fback_message='" + feedback_message + "'"
	query := "select pincode,city, state,district from PyApp_Postoffice"
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

	PostOfficeArr := []PostOfficeData{}

	var pincode int
	var state string
	var city string
	var district string

	for rows.Next() {
		rows.Scan(&pincode, &city, &state, &district)

		PostOfficeArr = append(PostOfficeArr, PostOfficeData{pincode, city, state, district})
	}
	ctx.JSON(200, PostOfficeArr)
}
