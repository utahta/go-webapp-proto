package session

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

const DefaultContextKey = "webapp_session_key"

var ErrorSessionNotFound = errors.New("Session not found")

type session struct {
	ctx     echo.Context
	name    string
	store   sessions.Store
	Session *sessions.Session
}

func New(c echo.Context, name string, store sessions.Store) {
	s := &session{ctx: c, name: name, store: store}
	c.Set(DefaultContextKey, s) // Context に持たせとく
}

func Start(c echo.Context) (*session, error) {
	s, ok := c.Get(DefaultContextKey).(*session)
	if !ok {
		return nil, ErrorSessionNotFound
	}

	ss, err := s.store.New(c.Request(), s.name)
	if err != nil {
		return nil, err
	}
	s.Session = ss
	return s, nil
}

func MustStart(c echo.Context) *session {
	s, err := Start(c)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// インターフェースは色々と検討の余地あり
// GetString とか？ interface{} オンリーは不便ぽさを感じる

func (s *session) Set(key string, v interface{}) {
	s.Session.Values[key] = v
}

func (s *session) Get(key string) (interface{}, bool) {
	v, ok := s.Session.Values[key]
	return v, ok
}

func (s *session) Save() error {
	return s.Session.Save(s.ctx.Request(), s.ctx.Response().Writer())
}
