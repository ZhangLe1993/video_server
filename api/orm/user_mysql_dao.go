package orm

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUser(loginName string, password string) (errs error) {

	stmt, err := conn.Prepare("INSERT INTO USER (login_name, password) VALUES (?, ?);")
	if err != nil {
		log.Printf("添加失败，失败信息：%s", err)
		return err
	}
	_, err = stmt.Exec(loginName, password)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func GetUser(loginName string) (rs string, errs error) {

	stmt, err := conn.Prepare("SELECT password FROM USER WHERE login_name = ?;")
	if err != nil {
		log.Printf("查询失败，失败信息：%s", err)
		return "", err
	}

	var password string
	err = stmt.QueryRow(loginName).Scan(&password)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmt.Close()
	return password, nil
}

func DeleteUser(loginName string, password string) (errs error) {
	stmt, err := conn.Prepare("DELETE FROM USER WHERE login_name = ? AND password = ? ;")
	if err != nil {
		log.Printf("删除失败，失败信息：%s", err)
		return err
	}
	_, err = stmt.Exec(loginName, password)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
