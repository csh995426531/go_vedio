package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempVid string

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
	err := AddUserCredential("崔赛航", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("崔赛航")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDelUser(t *testing.T) {
	err := DeleteUserCredential("崔赛航", "123")
	if err != nil {
		t.Errorf("Error of DelUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("崔赛航")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != "" {
		t.Error("Deleting user fail")
	}
}

func TestVedioWorkFlow(t *testing.T) {
	t.Run("AddVedio", testAddVedio)
	t.Run("GetVedio", testGetVedio)
	t.Run("DelVedio", testDelVedio)
	t.Run("RegetVedio", testRegetVedio)
}

func testAddVedio(t *testing.T) {
	info, err := AddNewVedio(1, "test")

	if err != nil {
		t.Errorf("Error of AddVedio : %v", err)
	}

	if info.Name != "test" {
		t.Error("Error of AddVedio Name Error")
	}
	tempVid = info.Id
}

func testGetVedio(t *testing.T) {
	info, err := GetVedio(tempVid)
	if err != nil {
		t.Errorf("Error of GetVedio : %v", err)
	}

	if info.Name != "test" {
		t.Error("Error of GetVedio Name Error")
	}
}

func testDelVedio(t *testing.T) {
	err := DelVedio(tempVid)
	if err != nil {
		t.Errorf("Error of DelVedio : %v", err)
	}
}

func testRegetVedio(t *testing.T) {
	_, err := GetVedio(tempVid)
	if err != nil {
		t.Errorf("Error of RegetVedio : %v", err)
	}
}

func TestComment(t *testing.T) {
	truncateTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComment", testAddComment)
	t.Run("ListComment", testListComment)
}

func testAddComment(t *testing.T) {
	err := AddComments("1", 1, "视频内容挺精彩的")
	if err != nil {
		t.Errorf("Error of AddComment : %v", err)
	}
	AddComments("1", 1, "视频内容挺无聊的")
}

func testListComment(t *testing.T) {
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	from := to - 86400
	res, err := ListComments("1", from, to)
	if err != nil {
		t.Errorf("Error of ListComment : %v", err)
	}
	fmt.Printf("comments: %v \n", res)
	for k, vv := range res {
		fmt.Printf("comment: %d, %v \n", k, vv)
	}
}
