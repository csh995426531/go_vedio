package session

import (
	"sync"
	"time"

	"github.com/csh995426531/go_vedio/api/dbops"
	"github.com/csh995426531/go_vedio/api/defs"
	"github.com/csh995426531/go_vedio/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

// LoadSessionsFromDB 从DB加载sessions
func LoadSessionsFromDB() {
	res, err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}
	res.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

//GenerateNewSessionID 生成新的session
func GenerateNewSessionID(username string) (string, error) {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000

	temp := &defs.SimpleSession{Username: username, TTL: ttl}
	sessionMap.Store(id, temp)
	if err := dbops.InsertSession(id, ttl, username); err != nil {
		return "", err
	}
	return id, nil
}

//IsSessionExpired 是否过期
func IsSessionExpired(sid string) (string, bool) {
	res, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ct > res.(*defs.SimpleSession).TTL {
			DeleteSession(sid)
			return "", true
		}
		return res.(*defs.SimpleSession).Username, false

	} else {
		return "", true
	}
}

// DeleteSession 删除session
func DeleteSession(sid string) {
	dbops.DeleteSession(sid)
	sessionMap.Delete(sid)
}
