package query

import "github.com/xavier-niu/bupt-queryer/request"

type AppQueryer struct {
	Session *request.Session
}

func NewAppQuery(session *request.Session) AppQueryer {
	return AppQueryer{
		Session: session,
	}
}
