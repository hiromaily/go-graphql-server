package main

import (
	"net/http"
	"testing"
)

type ResponseUser struct {
	Data struct {
		User User `json:"user"`
	} `json:"data"`
}

type User struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Country string `json:"country,omitempty"`
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
					Data: struct {
						User User `json:"user"`
					}{
						User{
							ID: 1,
						},
					},
				},
				err: nil,
			},
		},
	}

}
