package orm

import "testing"

var tempVid string

func TestVideoWorkFlow(t *testing.T) {
	t.Run("Add", testAddVideo)
	t.Run("Get", testGetVideoInfo)
	t.Run("Del", testDeleteVideo)
	t.Run("Reget", testReGetVideoInfo)
}

func testAddVideo(t *testing.T) {
	video, err := AddVideo(1, "测试视频-1")
	if err != nil {
		t.Errorf("添加视频出错，错误信息：%v", err)
	}
	tempVid = video.Id
}

func testGetVideoInfo(t *testing.T) {

	_, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("查询视频出错，错误信息：%v", err)
	}
}

func testDeleteVideo(t *testing.T) {

	err := DeleteVideo(tempVid)
	if err != nil {
		t.Errorf("删除视频出错，错误信息：%v", err)
	}
}

func testReGetVideoInfo(t *testing.T) {

	video, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("再次查询视频出错，错误信息：%v", err)
	}
	if video != nil {
		t.Errorf("删除视频出错，错误信息：%v", err)
	}
}
