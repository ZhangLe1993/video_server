package orm

import (
	"video_server/api/defs"
	"video_server/api/utils"
)

func AddNewComments(vid string, aid int, content string) error {

	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmt, err := conn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func GetCommentList(vid string, from, to int) (list []*defs.Comment, errs error) {

	stmt, err := conn.Prepare("SELECT a.id,b.login_name,a.content FROM comments a INNER JOIN user b ON a.author_id = b.id WHERE a.video_id = ? AND a.create_time > FROM_UNIXTIME(?) AND a.create_time <= FROM_UNIXTIME(?); ")
	var res []*defs.Comment
	if err != nil {
		return res, err
	}
	rows, err := stmt.Query(vid, from, to)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var id, author, content string
		if err := rows.Scan(&id, &author, &content); err != nil {
			return res, err
		}
		comment := &defs.Comment{id, vid, author, content}
		res = append(res, comment)
	}
	defer stmt.Close()
	/*for i, obj := range res {
		log.Printf("结果：%d, %v \n", i, obj)
		//fmt.Printf("结果：%d, %v \n", i, obj)
	}*/
	return res, nil
}
