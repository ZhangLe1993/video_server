package orm

import "testing"

// init(dblogin,truncate tables ) -> run tests -> clear data(truncate tables)

func TestUserWorkFlow(t *testing.T) {

	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testReGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUser("yule.zhang", "123456")
	if err != nil {
		t.Errorf("添加用户出错，错误信息：%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUser("yule.zhang")
	if pwd != "123456" || err != nil {
		t.Errorf("查询用户出错，错误信息：%v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("yule.zhang", "123456")
	if err != nil {
		t.Errorf("删除用户出错，错误信息：%v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUser("yule.zhang")
	if err != nil {
		t.Errorf("再次查询用户出错，错误信息：%v", err)
	}
	if pwd != "" {
		t.Errorf("删除用户出错，错误信息：%v", err)
	}
}
