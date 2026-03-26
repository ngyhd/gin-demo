package pkg

import (
	"reflect"
	"testing"
)

func TestSuccess(t *testing.T) {
	result := Success()

	if result["code"] != 0 {
		t.Errorf("Success() code = %v, want 0", result["code"])
	}
	if result["msg"] != "Success" {
		t.Errorf("Success() msg = %v, want Success", result["msg"])
	}
	if result["data"] != "" {
		t.Errorf("Success() data = %v, want empty string", result["data"])
	}
}

func TestSuccessWithData(t *testing.T) {
	tests := []struct {
		name string
		data any
	}{
		{
			name: "字符串数据",
			data: "test data",
		},
		{
			name: "数字数据",
			data: 123,
		},
		{
			name: "对象数据",
			data: map[string]any{"key": "value"},
		},
		{
			name: "nil数据",
			data: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SuccessWithData(tt.data)
			if result["code"] != 0 {
				t.Errorf("SuccessWithData() code = %v, want 0", result["code"])
			}
			if result["msg"] != "Success" {
				t.Errorf("SuccessWithData() msg = %v, want Success", result["msg"])
			}
			if !reflect.DeepEqual(result["data"], tt.data) {
				t.Errorf("SuccessWithData() data = %v, want %v", result["data"], tt.data)
			}
		})
	}
}

func TestFail(t *testing.T) {
	tests := []struct {
		name string
		code Code
		wantMsg string
	}{
		{
			name:    "参数错误",
			code:    ParamsErrCode,
			wantMsg: "参数错误",
		},
		{
			name:    "记录不存在",
			code:    RecordNotFoundErrCode,
			wantMsg: "记录不存在",
		},
		{
			name:    "用户已存在",
			code:    UserExistsErrCode,
			wantMsg: "用户已经存在",
		},
		{
			name:    "登录信息错误",
			code:    UserTokenErrCode,
			wantMsg: "登录信息错误",
		},
		{
			name:    "密码错误",
			code:    UserPasswordErrCode,
			wantMsg: "密码错误",
		},
		{
			name:    "邮箱已存在",
			code:    UserEmailExistsErrCode,
			wantMsg: "邮箱已经存在",
		},
		{
			name:    "系统内部错误",
			code:    InternalErrCode,
			wantMsg: "系统内部发生错误",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fail(tt.code)
			code := result["code"].(Code)
			if code != tt.code {
				t.Errorf("Fail() code = %v, want %v", code, tt.code)
			}
			if result["msg"] != tt.wantMsg {
				t.Errorf("Fail() msg = %v, want %v", result["msg"], tt.wantMsg)
			}
			if result["data"] != "" {
				t.Errorf("Fail() data = %v, want empty string", result["data"])
			}
		})
	}
}

func TestFailWithMessage(t *testing.T) {
	tests := []struct {
		name      string
		code      Code
		errMessage string
	}{
		{
			name:      "参数错误带消息",
			code:      ParamsErrCode,
			errMessage: "username is required",
		},
		{
			name:      "系统错误带消息",
			code:      InternalErrCode,
			errMessage: "database connection failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FailWithMessage(tt.code, tt.errMessage)
			expectedMsg := message[tt.code] + "." + tt.errMessage
			code := result["code"].(Code)
			if code != tt.code {
				t.Errorf("FailWithMessage() code = %v, want %v", code, tt.code)
			}
			if result["msg"] != expectedMsg {
				t.Errorf("FailWithMessage() msg = %v, want %v", result["msg"], expectedMsg)
			}
		})
	}
}

