package main

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "five",
			args: args{
				head: LinkedListFromArgs(1, 2, 3, 4, 5),
			},
			want: LinkedListFromArgs(3, 4, 5),
		},
		{
			name: "one",
			args: args{
				head: LinkedListFromArgs(1),
			},
			want: LinkedListFromArgs(1),
		},
		{
			name: "two",
			args: args{
				head: LinkedListFromArgs(1, 2),
			},
			want: LinkedListFromArgs(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
