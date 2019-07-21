package dbops

import "log"

// AddUserCredential 添加用户
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name, pwd) VALUES(?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

// GetUserCredential 获取用户
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FORM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

//DeleteUserCredential 删除用户
func DeleteUserCredential(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FORM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}
