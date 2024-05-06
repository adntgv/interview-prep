package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				head: LinkedListFromArgs(),
			},
			want: true,
		}, {
			name: "one",
			args: args{
				head: LinkedListFromArgs(1),
			},
			want: true,
		}, {
			name: "two",
			args: args{
				head: LinkedListFromArgs(1, 1),
			},
			want: true,
		}, {
			name: "three",
			args: args{
				head: LinkedListFromArgs(1, 2, 1),
			},
			want: true,
		},
		{
			name: "fail",
			args: args{
				head: LinkedListFromArgs(1, 2, 3),
			},
			want: true,
		},
		{
			name: "fail 2",
			args: args{
				head: LinkedListFromArgs(1, 2, 3, 1),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
