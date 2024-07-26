package auth

import "testing"

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		notWant string
		wantErr bool
	}{
		{
			name:    "basic case",
			args:    "password",
			notWant: "",
			wantErr: false,
		},
		{
			name:    "basic case",
			args:    "password",
			notWant: "password",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.notWant {
				t.Errorf("HashPassword() got = %v, notWant %v", got, tt.notWant)
			}
		})
	}
}

func TestCorrectPassword(t *testing.T) {
	tests := []struct {
		name      string
		plaintext string
		password  string
		want      bool
	}{
		{
			name:      "basic case",
			plaintext: "qwerty",
			password:  "qwerty",
			want:      true,
		},
		{
			name:      "wrong password case",
			plaintext: "wrong-password",
			password:  "password",
			want:      false,
		},
		{
			name:      "wrong password case",
			plaintext: "invalid-password",
			password:  "some-real-password",
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, _ := HashPassword(tt.password)
			if got := CorrectPassword(hashedPassword, []byte(tt.plaintext)); got != tt.want {
				t.Errorf("CorrectPassword() = %v, notWant %v", got, tt.want)
			}
		})
	}
}
