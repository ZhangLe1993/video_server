package orm

import (
	"database/sql"
	"strconv"
	"sync"
	"video_server/api/defs"
)

func InsertSession(sid string, ttl int64, name string) (errs error) {

	ttlStr := strconv.FormatInt(ttl, 10)

	stmt, err := conn.Prepare("INSERT INTO sessions (session_id,login_name,TTL) VALUES (?,?,?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sid, ttlStr, name)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func GetSession(sid string) (session *defs.SessionVO, errs error) {

	stmt, err := conn.Prepare("SeLECT login_name,TTL FROM sessions WHERE　session_id = ? limit 0,1;")
	if err != nil {
		return nil, err
	}
	var name, ttl string
	err = stmt.QueryRow(sid).Scan(&name, &ttl)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	vo := &defs.SessionVO{}

	if res, err := strconv.ParseInt(ttl, 64, 10); err == nil {
		vo.UserName = name
		vo.TTL = res
	} else {
		return nil, err
	}

	defer stmt.Close()

	return vo, nil
}

func GetAllSession() (sessionMap *sync.Map, errs error) {

	stmt, err := conn.Prepare("SELECT session_id,login_name,TTL FROM sessions;")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	//定义一个Map
	cacheMap := &sync.Map{}
	for rows.Next() {
		var sessionId string
		var name string
		var ttlTemp string
		if err = rows.Scan(&sessionId, &name, &ttlTemp); err != nil {
			return nil, err
		}
		if ttl, err := strconv.ParseInt(ttlTemp, 64, 10); err == nil {
			//获取session
			session := &defs.SessionVO{name, ttl}
			//存入Map
			cacheMap.Store(sessionId, session)
		}
	}
	defer stmt.Close()
	return cacheMap, nil
}

func DeleteSession(sid string) error {
	stmt, err := conn.Prepare("DELETE FROM　sessions　WHERE session_id = ?; ")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sid)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
