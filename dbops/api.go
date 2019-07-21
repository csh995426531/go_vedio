package dbops

import (
	"database/sql"
	"log"
	"time"
	"vedio/defs"
	"vedio/utils"
)

// AddUserCredential 添加用户
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name, pwd) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// GetUserCredential 获取用户
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s", err)
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

//DeleteUserCredential 删除用户
func DeleteUserCredential(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

// AddNewVedio 新增视频
func AddNewVedio(aid int, name string) (*defs.VedioInfo, error) {

	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	time := time.Now()
	ctime := time.Format("Jan 02 2006 15:04:05")
	stmtIns, err := dbConn.Prepare("INSERT INTO vedio_info (id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	res := &defs.VedioInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	return res, nil
}

//GetVedio 获取视频
func GetVedio(vid string) (*defs.VedioInfo, error) {

	stmtOut, err := dbConn.Prepare("SELECT author_id,name,display_ctime FROM vedio_info WHERE id = ?")
	if err != nil {
		return nil, err
	}

	var aid int
	var name string
	var ctime string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	defer stmtOut.Close()

	res := &defs.VedioInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	return res, nil
}

//DelVedio 删除视频
func DelVedio(vid string) error {

	stmtDel, err := dbConn.Prepare("DELETE FROM vedio_info WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}
