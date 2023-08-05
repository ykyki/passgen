package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "option -v",
			args: []string{"-v"},
			want: 0,
		},
		{
			name: "option --version",
			args: []string{"--version"},
			want: 0,
		},
		{
			name: "option --length",
			args: []string{"--length", "32"},
			want: 0,
		},
		{
			name: "option -l",
			args: []string{"-l", "8"},
			want: 0,
		},
		{
			name: "option -l=4",
			args: []string{"-l=4"},
			want: 0,
		},
		{
			name: "option -l with negative value",
			args: []string{"-l='-1'"},
			want: 1,
		},
		{
			name: "option -l with negative value",
			args: []string{"--length '-2'"},
			want: 1,
		},
		{
			name: "option undefined",
			args: []string{"-q"},
			want: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cli := &cli{outStream: new(bytes.Buffer), errStream: new(bytes.Buffer)}

			got := cli.run(append([]string{"passgen"}, tc.args...))

			if got != tc.want {
				t.Errorf("cli.run(%v) = %v, want %v", tc.args, got, tc.want)
			}
		})
	}
}
