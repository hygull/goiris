package conf

type CustomError struct {
	Type       string `json:"type,omitempty"`
	Details    string `json:"details,omitempty"`
	StatusCode int    `json:"status code,omitempty"`
}

var CustomErr = []CustomError{
	{"InsertQueryExecutionError", "Error in execution of insert query", 101},
	{"SelectQueryExecutionError", "Error in execution of select query", 102},
	{"UpdateQueryExecutionError", "Error in execution of update query", 103},
	{"DeleteQueryExecutionError", "Error in execution of delete query", 104},
	{"UnequalKeysForValuesError", "The number of keys and their related values " +
		"should be equal", 105},
	{"DatabaseConnectionError", "Error while connecting to database, please check credentials", 106},
	{"EmptyFieldOrWrongKeyNameError", "You are not providing all the key's value or you've provided wrong key names", 107},
	{"WrongQueryIndexError", "The query index should range from 1 to 4...(insert, select, update delete)", 108},
	{"UserDoesNotExistError", "This user id does not exist in the database", 109},
}
