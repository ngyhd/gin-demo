package pkg

import (
	"testing"
)

var (
	testHashCorrect = HashPassword("123456")
	testHashOther   = HashPassword("otherpass")
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantLen  int
	}{
		{
			name:     "正常密码",
			password: "123456",
			wantLen:  60,
		},
		{
			name:     "空密码",
			password: "",
			wantLen:  60,
		},
		{
			name:     "长密码",
			password: "abcdefghijklmnopqrstuvwxyz",
			wantLen:  60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HashPassword(tt.password)
			if len(got) != tt.wantLen {
				t.Errorf("HashPassword() length = %v, want %v", len(got), tt.wantLen)
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	tests := []struct {
		name          string
		password      string
		hashedPassword string
		wantErr       bool
	}{
		{
			name:          "正确密码",
			password:      "123456",
			hashedPassword: testHashCorrect,
			wantErr:       false,
		},
		{
			name:          "错误密码",
			password:      "123456",
			hashedPassword: testHashOther,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPassword(tt.hashedPassword, tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashAndCheckPassword(t *testing.T) {
	password := "testPassword123"
	hashed := HashPassword(password)

	err := CheckPassword(hashed, password)
	if err != nil {
		t.Errorf("CheckPassword() failed with correct password: %v", err)
	}

	err = CheckPassword(hashed, "wrongPassword")
	if err == nil {
		t.Errorf("CheckPassword() should fail with wrong password")
	}
}

