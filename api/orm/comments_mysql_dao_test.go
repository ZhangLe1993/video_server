package orm

import (
	"strconv"
	"testing"
	"time"
)

func TestCommentsWorkFlow(t *testing.T) {
	//t.Run("AddUSER",testAddUser)
	t.Run("AddCOMMENT", testAddNewComments)
	t.Run("GetCOMMENT", testGetCommentList)
}

func testAddNewComments(t *testing.T) {

	err := AddNewComments("111", 1, "这不是演习")
	if err != nil {
		t.Errorf("添加评论出错，错误信息：%v", err)
	}

}

func testGetCommentList(t *testing.T) {

	var vid = "111"
	var from = 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := GetCommentList(vid, from, to)
	if err != nil {
		t.Errorf("查询评论出错，错误信息：%v", err)
	}
	for i, obj := range res {
		t.Logf("结果：%d, %v \n", i, obj)
	}

}
