/*      ############  JSON  VIEW  ###################

Aim 	   : To display the JSON data to the browser
Created on : 18/10/2016
*/

package views

import (
	"fmt"
	"github.com/kataras/iris"
)
import (
	"goiris/conf"
)

/************* Construction of JSON responsers for requests *********************/

//This function will display the required JSON Response for a particular Request
func ShowJSON(ctx *iris.Context, keysSlice []string, valuesSlice []interface{}) {
	if lengthOfKeysAndValuesSlicesAreEqual(len(keysSlice), len(valuesSlice)) {
		jsonDict := iris.Map{}
		i := 0
		for _, key := range keysSlice {
			jsonDict[key] = valuesSlice[i]
			i++
		}
		ctx.JSON(iris.StatusOK, jsonDict)
	} else {
		fmt.Println("\n[SynkkU] The number of keys(strings) and values should be same in both slices")
		ShowErrorAsJSON(ctx, 5)
	}
}

func lengthOfKeysAndValuesSlicesAreEqual(keyLen int, vLen int) bool {
	return keyLen == vLen
}

func ShowErrorAsJSON(ctx *iris.Context, index int) {
	ctx.JSON(200, iris.Map{"Error": conf.CustomErr[index-1]})
}
