package session

import (
	"sync"
	"time"
	"video_server/api/defs"
	"video_server/api/orm"
	"video_server/api/utils"
)

var sessionCache *sync.Map

func init() {
	sessionCache = &sync.Map{}
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteSession(sid string) {
	orm.DeleteSession(sid)
	sessionCache.Delete(sid)
}

func LoadSessionFromDB() {
	res, err := orm.GetAllSession()
	if err != nil {
		return
	}
	res.Range(func(key, value interface{}) bool {
		temp := value.(*defs.SessionVO)
		sessionCache.Store(key, temp)
		return true
	})
}

func GenerateSessionId(name string) (sessionId string) {
	sid, err := utils.NewUUID()
	if err != nil {
		return ""
	}
	//创建时间，毫秒
	currentTime := currentTimeMillis()
	//过期时间，30分钟
	ttl := currentTime + 30*60*1000

	session := &defs.SessionVO{name, ttl}
	sessionCache.Store(sid, session)
	orm.InsertSession(sid, ttl, name)
	return sid
}

func IsSessionExpire(sessionId string) (sid string, expire bool) {

	session, ok := sessionCache.Load(sessionId)
	if ok {
		currentTime := currentTimeMillis()
		//ttl = 存储时的系统毫秒数  + 30的毫秒数
		if session.(*defs.SessionVO).TTL < currentTime {
			//删除过期session
			deleteSession(sid)
			return "", true
		}
		return session.(*defs.SessionVO).UserName, false
	}
	return "", true
}
