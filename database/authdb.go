package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/xoraes/dappauth/models"
	"log"
	"os"
)
var DB *sql.DB
func SetDBPool(db *sql.DB) {
	DB = db
}
func init() {
	// load .env file
	hostport := os.Getenv("DBHOST") + ":" + os.Getenv("PORT")
	dbname := os.Getenv("POSTGRES_DB")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	sslmode := os.Getenv("SSLMODE")

	connUrl := "postgresql://" + hostport + "/" + dbname + "?user=" +
		username + "&password=" + password + "&sslmode=" + sslmode
	// Open the connection
	db, err := sql.Open("postgres", connUrl)
	if err != nil {
		panic(err)
	}
	// check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	SetDBPool(db)
	logrus.Info("Successfully connected to DB at:", hostport)
}
func InsertUser(email, hashPassword, firstName, lastName string) error {

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO users (email, hashpassword, firstname, lastname) VALUES ($1, $2, $3, $4)`
	// execute the sql statement
	err := DB.QueryRow(sqlStatement, email, hashPassword, firstName, lastName)

	if err.Err() != nil {
		return err.Err()
	}
	logrus.Debugf("Inserted a single record %v", email)
	// return the inserted id
	return nil
}
func GetUser(email string) (*models.User, error) {
	// create a user of models.User type
	var user models.User
	// create the select sql query
	sqlStatement := `SELECT email, hashpassword FROM users WHERE email=$1`
	// execute the sql statement
	row := DB.QueryRow(sqlStatement, email)
	// unmarshal the row object to user
	err := row.Scan(&user.Email, &user.Password)
	switch err {
	case sql.ErrNoRows:
		logrus.Debugf("No rows were returned!")
		return &user, nil
	case nil:
		return &user, nil
	default:
		logrus.Debugf("Unable to scan the row. %v", err)
	}
	// return empty user on error
	return &user, err
}
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	// create the select sql query
	sqlStatement := `SELECT email, firstname, lastname FROM users`
	// execute the sql statement
	rows, err := DB.Query(sqlStatement)
	if err != nil {
		logrus.Debugf("Unable to execute the query. %v", err)
	}
	// close the statement
	defer rows.Close()
	// iterate over the rows
	for rows.Next() {
		var user models.User
		// unmarshal the row object to user
		err = rows.Scan(&user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// append the user in the users slice
		users = append(users, user)
	}
	return users, err
}
func UpdateUser(email, firstName, lastName string) (int64, error) {
	if firstName == "" && lastName == "" {
		return 0, nil
	}
	// create the update sql query
	var sqlStatement string
	var res sql.Result
	var err error
	if firstName == "" {
		sqlStatement = `UPDATE users SET lastname=$2 WHERE email=$1`
		res, err = DB.Exec(sqlStatement, email, lastName)
	} else if lastName == "" {
		sqlStatement = `UPDATE users SET firstname=$2 WHERE email=$1`
		res, err = DB.Exec(sqlStatement, email, firstName)
	} else {
		sqlStatement = `UPDATE users SET firstname=$2, lastName=$3 WHERE email=$1`
		res, err = DB.Exec(sqlStatement, email, firstName, lastName)
	}
	// execute the sql statement
	if err != nil {
		logrus.Debugf("Unable to execute the query. %v", err)
		return 0, err
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.Debugf("Error while checking the affected rows. %v", err)
		return 0, err
	}
	logrus.Debugf("Total rows/record affected %v", rowsAffected)
	return rowsAffected, nil
}
