package main

import (
	"bytes"
	"encoding/json"
	"gin-demo/internal"
	"github.com/gin-gonic/gin"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
)

// postJson 以Json形式传递参数，发起post请求
func postJson(uri string, param map[string]interface{}, router *gin.Engine) *httptest.ResponseRecorder {
	jsonByte, _ := json.Marshal(param)
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
	req.Header.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.nKqRzljFfJKlotnxH8auq7ui3jlIZVxI16VZQ0G0yVY")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// 测试API
func TestApi(t *testing.T) {
	tests := []struct {
		name   string
		uri    string
		params map[string]any
	}{
		{
			name: "注册",
			uri:  "/register",
			params: map[string]any{
				"username": "Sean2",
				"password": "123456",
				"email":    "8888888@qq.com",
			},
		},
		{
			name: "登录",
			uri:  "/login",
			params: map[string]any{
				"username": "Sean3",
				"password": "123456",
			},
		},
		{
			name: "用户信息",
			uri:  "/user/info",
			params: map[string]any{
				"id": 2,
			},
		},
		{
			name: "更新信息",
			uri:  "/user/update",
			params: map[string]any{
				"username": "Sean3",
				"id":       14,
			},
		},
		{
			name: "删除用户",
			uri:  "/user/delete",
			params: map[string]any{
				"id": 13,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := postJson(tt.uri, tt.params, internal.Exec())
			m := map[string]any{}
			data, err := io.ReadAll(w.Body)
			if err != nil {
				t.Errorf("请求错误, err:%s", err.Error())
			}
			err = json.Unmarshal(data, &m)
			if !reflect.DeepEqual(w.Code, 200) || !reflect.DeepEqual(err, nil) || !reflect.DeepEqual(m["code"], float64(0)) {
				t.Errorf("响应数据不符，errmsg:%v\n", m["msg"])
			}

		})
	}
}
