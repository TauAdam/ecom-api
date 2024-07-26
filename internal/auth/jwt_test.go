package auth

import "testing"

func TestCreateJWToken(t *testing.T) {
	type args struct {
		secret []byte
		userID int
	}
	tests := []struct {
		name    string
		args    args
		notWant string
		wantErr bool
	}{
		{
			name: "basic case",
			args: args{
				secret: []byte("secret"),
				userID: 1,
			},
			notWant: "",
			wantErr: false,
		},
		{
			name: "basic case",
			args: args{
				secret: []byte("secret"),
				userID: 1,
			},
			notWant: "some-wrong-token",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateJWToken(tt.args.secret, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJWToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.notWant {
				t.Errorf("CreateJWToken() got = %v, notWant %v", got, tt.notWant)
			}
		})
	}
}
