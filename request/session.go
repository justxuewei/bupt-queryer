package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	checkUrl = "https://app.bupt.edu.cn/uc/wap/login/check"
)

type Session struct {
	client *http.Client
	isLogin bool
	usr, pwd string
}

func NewSession(usr, pwd string) *Session {
	jar, _ := cookiejar.New(nil)
	session := &Session{
		client: &http.Client{
			Jar: jar,
		},
		usr: usr,
		pwd: pwd,
	}
	return session
}

func (s *Session) login() error {
	rsp, err := s.client.PostForm(checkUrl, url.Values{
		"username": {s.usr},
		"password": {s.pwd},
	})
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	var info BaseInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		return err
	}

	if info.Code != 0 {
		return errors.New(fmt.Sprintf("%s", info.Message))
	}
	s.isLogin = true
	return nil
}

func (s *Session) PostForm(url string, data url.Values) (*http.Response, error) {
	if !s.isLogin {
		if err := s.login(); err != nil {
			return nil, err
		}
	}
	return s.client.PostForm(url, data)
}
