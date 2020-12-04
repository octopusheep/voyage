package dbops

import(
	"testing"
)

// init(dblogin,truncate tables)->run tests -> clear data(truncate tables)

func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M){
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T){
	t.Run("Add",testAddUser)
	t.Run("Get",testgetUser)
	t.Run("Delete",testDeleteUser)
	t.Run("Reget",testRegetUser)

}

func testAddUser(t *testing.T){
	err:=AddUserCredential("zhangyuayng","123")
	if err!=nil{
		t.Errorf("Error of AddUser: %v",err)
	}
}

func testgetUser(t *testing.T){
	pwd,err:=GetUserCredential("zhangyuayng")
	if pwd!="123"||err!=nil{
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T){
	err:=DeleteUser("zhangyuayng","123")
	if err!=nil{
		t.Errorf("Error of DeleteUser: %v",err)
	}
}

func testRegetUser(t *testing.T){
	pwd,err:=GetUserCredential("zhangyuayng")
	if err!=nil{
		t.Errorf("Error of RegetUser: %v",err)
	}

	if pwd!=""{
		t.Errorf("Deleting user test failed")
	}
}
