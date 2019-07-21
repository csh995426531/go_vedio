package dbops

import "testing"

func truncateTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate vedio_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {

	truncateTables()
	m.Run()
	truncateTables()
}

func TestUserWorkFlow(t *testing.T) {

	t.Run("AddUser", testAddUser)
	t.Run("GetUser", testGetUser)
	t.Run("DelUser", testDelUser)
	t.Run("RegetUser", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("test", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDelUser(t *testing.T) {
	err := DeleteUserCredential("test", "123")
	if err != nil {
		t.Errorf("Error of DelUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if pwd != "" || err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
}
