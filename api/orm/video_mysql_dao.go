package orm

import (
	"database/sql"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

func AddVideo(authorId int, name string) (video *defs.Video, errs error) {

	videoId, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	// Jan 02 2006, 15:04:05 时间原点
	ctime := t.Format("Jan 02 2006, 15:04:05")

	stmt, err := conn.Prepare("INSERT INTO video_info (id,author_id,name,display_ctime) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(videoId, authorId, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.Video{videoId, authorId, name, ctime}

	defer stmt.Close()
	return res, nil
}

func GetVideoInfo(videoId string) (video *defs.Video, errs error) {
	stmt, err := conn.Prepare("SELECT author_id,name,display_ctime FROM video_info WHERE id = ?; ")

	if err != nil {
		return nil, err
	}
	var aid int
	var dct string
	var name string
	err = stmt.QueryRow(videoId).Scan(&aid, &name, &dct)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmt.Close()
	res := &defs.Video{videoId, aid, name, dct}

	return res, nil

}

func DeleteVideo(videoId string) (errs error) {
	stmt, err := conn.Prepare("DELETE FROM video_info WHERE id = ?; ")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(videoId)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
