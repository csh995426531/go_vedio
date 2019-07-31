package dbops

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/csh995426531/go_vedio/api/defs"
	"github.com/csh995426531/go_vedio/api/utils"
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
	res, err := stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	if delNum, _ := res.RowsAffected(); delNum != 1 {
		err := errors.New("This user delete fail")
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

// AddComments 增加评论
func AddComments(vid string, aid int, content string) error {
	cid, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, vedio_id, author_id, content) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(cid, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

// ListComments 评论列表
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.vedio_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	if err != nil {
		return res, err
	}

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string

		err := rows.Scan(&id, &name, &content)
		if err != nil {
			return res, err
		}

		temp := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, temp)
	}

	defer stmtOut.Close()
	return res, nil
}
