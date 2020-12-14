package models

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Fermekoo/blog-dandi/helpers"

	"github.com/Fermekoo/blog-dandi/db"
)

//User ..
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

//FetchAllUser Mengambil data user..
func FetchAllUser() (Response, error) {
	var obj User
	var arrobj []User
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id,email,fullname, password FROM df_users"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Email, &obj.Fullname, &obj.Password)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreUser(email string, fullname string, password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO df_users (email, password, fullname) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(email, password, fullname)

	if err != nil {
		return res, err
	}

	lastInsertedID, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "mantap gan"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedID,
	}

	return res, nil
}

func UpdateUser(id int, email string, fullname string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE df_users SET email = ?, fullname = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(email, fullname, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM df_users WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func CheckLogin(email, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT id, email, fullname, password FROM df_users WHERE email = ?"

	err := con.QueryRow(sqlStatement, email).Scan(
		&obj.ID, &obj.Email, &obj.Fullname, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("user not found")

		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)

	if !match {
		fmt.Println("password doesn't match")

		return false, err
	}

	return true, nil
}
