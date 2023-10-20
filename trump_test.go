package trump

import "testing"

func TestCovfefe(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "basic_tweet",
			input:  "Despite the negative press",
			output: "Despite the negative press covfefe",
		},
		{
			name:   "short_tweet",
			input:  "MAGA",
			output: "MAGA covfefe",
		},
		{
			name:   "long_tweet",
			input:  "Fake News Media is working overtime to blame me for the horrible attack in New Zealand",
			output: "Fake News Media is working overtime to blame me for the horrible attack in New Zealand covfefe",
		},
		{
			name:   "no_input",
			input:  "",
			output: " covfefe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Covfefe(tt.input)
			if got != tt.output {
				t.Errorf("For input %q, expected %q, but got %q", tt.input, tt.output, got)
			}
		})
	}
}
