package dbops

import(
	"testing"
	"strconv"
	"time"
	"fmt"
)

var tempvid string

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

func TestVideoWorkFlow(t *testing.T){
	clearTables()
	t.Run("PrepareUser",testAddUser)
	t.Run("AddVideo",testgetUser)
	t.Run("GetVideo",testgetUser)
	t.Run("DeleteVideo",testDeleteUser)
	t.Run("RegetVideo",testRegetUser)

}

func testAddNewVideo(t *testing.T){
	vi,err:=AddNewVideo(1,"my-video")
	if err!=nil{
		t.Errorf("Error of AddVideo: %v",err)
	}

	tempvid=vi.Id
}

func testgetVideoInfo(t *testing.T){
	_,err:=GetVideoInfo(tempvid)
	if err!=nil{
		t.Errorf("Error of getVideo:%v",err)
	}
}

func testDeleteVideoInfo(t *testing.T){
	err:=DeleteVideoInfo(tempvid)
	if err!=nil{
		t.Errorf("Error of DeleteVideoInfo: %v",err)
	}
}

func testRegetVideoInfo(t *testing.T){
	vi,err:=GetVideoInfo(tempvid)
	if err!=nil||vi!=nil{
		t.Errorf("Error of RegetVideo: %v",err)
	}
}

func TestCommentsWorkFlow(t *testing.T){
	clearTables()
	t.Run("PrepareUser",testAddUser)
	t.Run("AddComments",testAddComments)
	t.Run("ListComments",testListComments)
}

func testAddComments(t *testing.T){
	vid:="12345"
	aid:=1
	content:="I like this video"

	err:=AddNewComments(vid,aid,content)

	if err!= nil{
		t.Errorf("Error of AddComments: %v",err)
	}
}

func testListComments(t *testing.T){
	vid:="12345"
	from:=1514764800
	to,_:=strconv.Atoi(strconv.FromInt(time,Now().UnixNano()/1000000000,10))

	res,err:=ListComments(vid,from,to)
	if err!=nil{
		t.Errorf("Error of ListComments: %v",err)
	}

	for i,ele :=range res{
		fmt.Printf("comments: %d,%v \n",i,ele)
	}
	err:=AddNewComments(vid,aid,content)

	if err!= nil{
		t.Errorf("Error of AddComments: %v",err)
	}
}


