//Package name declaration
package conf

//Name of root directory inside $GOPATH/src/  (in MAC)    or  %GOPATH%\src\ (in windows)
//In my case it is 'goiris' ... and 'goiris' is a main root folder containg ...
//conf, controllers, routers, views, models, docs and other files ...
//If you want other names for root directory eg.  'synkku' or 'synkkuapi' or any other...
//name then you can set it here by assigning a new string value to RootDir variable
var RootDir = "goiris"

//Port number of web service on which it will listen for requests ...
var ServerPort = "8080"

/******************* Database credentials ******************************************/
//Name of the database used
var Db = "django"

//IP address of host on which database is running...You can also leave it a blank string
//to allow it to listen for any IP
var DbHost = "127.0.0.1"

//Port number for database server on which it is listening
var DbPort = "3306"

//Username for the database...root is default name in case of MySQL
var DbUser = "root"

//Password for the database
var DbPassword = "admin@321"

//Type of the database used...eg.  'mysql' or 'postgres'...
//In our case these both databases are being used...but we are preferring 'postgres'
var DbType = "mysql"
