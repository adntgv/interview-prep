package main

import (
	"testing"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case 1",
			args: args{
				head: LinkedListFromArgs(1, 2, 3, 4, 5),
			},
			want: LinkedListFromArgs(1, 5, 2, 4, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			got := tt.args.head
			if !listsEq(got, tt.want) {
				t.Errorf("reorderList() = %v, want %v", linkedListToString(got), linkedListToString(tt.want))
			}
		})
	}
}

func Test_listsEq(t *testing.T) {
	type args struct {
		first  *ListNode
		second *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal",
			args: args{
				first:  LinkedListFromArgs(1, 2, 3, 4, 5),
				second: LinkedListFromArgs(1, 2, 3, 4, 5),
			},
			want: true,
		},
		{
			name: "equal",
			args: args{
				first:  LinkedListFromArgs(1, 2, 3, 4, 5),
				second: LinkedListFromArgs(1, 2, 3, 4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listsEq(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("listsEq() = %v, want %v", got, tt.want)
			}
		})
	}
}
