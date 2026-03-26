package main

import (
	"bytes"
	"encoding/json"
	"gin-demo/internal"
	"github.com/gin-gonic/gin"
	"io"
	"net/http/httptest"
	"os"
	"testing"
)

const testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.nKqRzljFfJKlotnxH8auq7ui3jlIZVxI16VZQ0G0yVY"

// postJson 以Json形式传递参数，发起post请求
func postJson(uri string, param map[string]interface{}, router *gin.Engine) *httptest.ResponseRecorder {
	jsonByte, _ := json.Marshal(param)
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
	req.Header.Set("token", testToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// skipIfNoDatabase 在数据库不可用时跳过测试
func skipIfNoDatabase(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "1" {
		t.Skip("跳过集成测试（设置 SKIP_INTEGRATION_TESTS=1 启用）")
	}
	// 尝试验证数据库连接是否可用
	// 由于 internal.Exec() 会在 MySQL 不可用时 panic，
	// 我们通过环境变量来判断是否跳过
	if os.Getenv("MYSQL_HOST") == "" && os.Getenv("DB_HOST") == "" {
		t.Skip("跳过集成测试：未配置数据库（设置 SKIP_INTEGRATION_TESTS=1 可跳过此检查）")
	}
}

// TestApi 测试API（需要数据库支持）
func TestApi(t *testing.T) {
	skipIfNoDatabase(t)

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
			if w.Code != 200 || err != nil || m["code"] != float64(0) {
				t.Errorf("响应数据不符，errmsg:%v\n", m["msg"])
			}

		})
	}
}
