package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

type ResponseUserList struct {
	Data UserListData `json:"data"`
}

type UserListData struct {
	UserList []User `json:"userList"`
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
			},
		},
		{
			name: "user for id, name",
			args: args{
				url:     `/graphql?query={user(id:"1"){id,name}}`,
				method:  "GET",
				headers: nil,
			},
			want: want{
				statusCode: http.StatusOK,
				respUser: &ResponseUser{
					Data: UserData{
						User{
							ID:   1,
							Name: "Dan",
						},
					},
				},
			},
		},
		{
			name: "user for id, name, age",
			args: args{
				url:     `/graphql?query={user(id:"1"){id,name,age}}`,
				method:  "GET",
				headers: nil,
			},
			want: want{
				statusCode: http.StatusOK,
				respUser: &ResponseUser{
					Data: UserData{
						User{
							ID:   1,
							Name: "Dan",
							Age:  24,
						},
					},
				},
			},
		},
		{
			name: "user for id, name, age, country",
			args: args{
				url:     `/graphql?query={user(id:"1"){id,name,age,country}}`,
				method:  "GET",
				headers: nil,
			},
			want: want{
				statusCode: http.StatusOK,
				respUser: &ResponseUser{
					Data: UserData{
						User{
							ID:      1,
							Name:    "Dan",
							Age:     24,
							Country: "United States",
						},
					},
				},
			},
		},
		// TODO: add resume
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
			if res.StatusCode != tt.want.statusCode {
				t.Errorf("status code: got %d, but want %d", res.StatusCode, tt.want.statusCode)
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
			// debug.DigIn(tt.want.respUser)
			if !reflect.DeepEqual(respUser, *tt.want.respUser) {
				t.Errorf("[url: %s] got = %v, want %v", tt.args.url, respUser, tt.want.respUser)
			}
		})
	}
}

func TestQueryUserList(t *testing.T) {
	type args struct {
		url     string
		method  string
		headers []map[string]string
	}
	type want struct {
		statusCode   int
		respUserList *ResponseUserList
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "userList for id",
			args: args{
				url:     `/graphql?query={userList{id,name}}`,
				method:  "GET",
				headers: nil,
			},
			want: want{
				statusCode: http.StatusOK,
				respUserList: &ResponseUserList{
					Data: UserListData{
						UserList: []User{
							{
								ID:   1,
								Name: "Dan",
							},
							{
								ID:   2,
								Name: "Lee",
							},
							{
								ID:   3,
								Name: "Nick",
							},
						},
					},
				},
			},
		},
		// TODO: add resume
	}
	ts, err := runTestServer()
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()
	client := getClient()

	opts := []cmp.Option{
		cmpopts.SortMaps(func(x, y int) bool { return x < y }),
		cmpopts.SortSlices(func(x, y User) bool { return x.ID < y.ID }),
	}

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
			if res.StatusCode != tt.want.statusCode {
				t.Errorf("status code: got %d, but want %d", res.StatusCode, tt.want.statusCode)
				return
			}
			// check body
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error("fail to call io.ReadAll(res.Body)")
				return
			}
			var respUserList ResponseUserList
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			err = json.Unmarshal(body, &respUserList)

			debug.DigIn(respUserList)
			// Note: order in list is not guaranteed
			if !cmp.Equal(respUserList, *tt.want.respUserList, opts...) {
				// if !reflect.DeepEqual(respUserList, *tt.want.respUserList) {
				t.Errorf("[url: %s] got = %v, want %v", tt.args.url, respUserList, tt.want.respUserList)
			}
		})
	}
}
