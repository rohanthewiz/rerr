package rerr

import "testing"

func TestLastNTokens(t *testing.T) {
	type args struct {
		str       string
		separator string
		numTokens int
	}
	tests := []struct {
		name           string
		args           args
		wantLastTokens string
	}{
		{
			name:           "Test Last N Tokens 1",
			args:           args{str: "abc/def/ghi", separator: "/", numTokens: 2},
			wantLastTokens: "def/ghi",
		},
		{
			name:           "Test Last N Tokens 2",
			args:           args{str: "abcdefg", separator: "/", numTokens: 2},
			wantLastTokens: "abcdefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLastTokens := LastNTokens(tt.args.str, tt.args.separator, tt.args.numTokens); gotLastTokens != tt.wantLastTokens {
				t.Errorf("Last2Tokens() = %v, want %v", gotLastTokens, tt.wantLastTokens)
			}
		})
	}
}
