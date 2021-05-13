package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"github.com/hiromaily/go-graphql-server/pkg/debug"
)

type ResponseUser struct {
	Data UserData `json:"data"`
}

type UserData struct {
	User User `json:"user"`
}

type User struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Country string `json:"country,omitempty"`
}

func runTestServer() (*httptest.Server, error) {
	conf := getConfig()
	regi := NewRegistry(conf)

	srv := regi.NewServer()
	return srv.StartTest()
}

func setHTTPHeaders(req *http.Request, headers []map[string]string) {
	// req.Header.Set("Authorization", "Bearer access-token")
	for _, header := range headers {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
}

func getClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(3) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("redirect")
		},
	}
}

func TestQueryUser(t *testing.T) {
	type args struct {
		url     string
		method  string
		headers []map[string]string
	}
	type want struct {
		statusCode int
		respUser   *ResponseUser
		err        error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "user for id",
			args: args{
				url:     `/graphql?query={user(id:"1"){id}}`,
				method:  "GET",
				headers: nil,
			},
			want: want{
				statusCode: http.StatusOK,
				respUser: &ResponseUser{
					Data: UserData{
						User{
							ID: 1,
						},
					},
				},
				err: nil,
			},
		},
	}
	ts, err := runTestServer()
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()
	client := getClient()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.args.method, fmt.Sprintf("%s%s", ts.URL, tt.args.url), nil)
			if tt.args.headers != nil {
				setHTTPHeaders(req, tt.args.headers)
			}
			// request
			res, err := client.Do(req)
			defer func() {
				if res.Body != nil {
					res.Body.Close()
				}
			}()
			if err != nil {
				t.Errorf("fail to call [%s]", tt.args.url)
				return
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("status code is not 200, %d was returned", res.StatusCode)
				return
			}
			// check body
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error("fail to call io.ReadAll(res.Body)")
				return
			}
			var respUser ResponseUser
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			err = json.Unmarshal(body, &respUser)

			debug.DigIn(respUser)
		})
	}
}
