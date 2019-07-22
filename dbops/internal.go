package dbops

import (
	"strconv"
	"sync"
	"vedio/defs"
)

// InsertSession 写入session
func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions(session_id, TTL, login_name) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// RetrieveSession 取出session
func RetrieveSession(sid string) (*defs.SimpleSession, error) {

	res := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return res, err
	}
	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil {
		return res, err
	}

	defer stmtOut.Close()

	if temp, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		res.TTL = temp
		res.Username = uname
	} else {
		return res, err
	}

	return res, nil
}

// RetrieveAllSession 取出全部session
func RetrieveAllSession() (*sync.Map, error) {
	res := &sync.Map{}

	stmtOut, err := dbConn.Prepare("SELECT session_id, TTL, login_name FROM sessions")
	if err != nil {
		return res, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		return res, err
	}

	defer stmtOut.Close()

	for rows.Next() {
		var id string
		var ttlstr string
		var uname string
		if err = rows.Scan(&id, &ttlstr, &uname); err != nil {
			return res, err
		}

		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			temp := &defs.SimpleSession{Username: uname, TTL: ttl}
			res.Store(id, temp)
		}
	}
	return res, nil
}

// DeleteSession 删除session
func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM session WHERE session_id = ?")
	if err != nil {
		return err
	}
	if _, err = stmtDel.Exec(sid); err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}
