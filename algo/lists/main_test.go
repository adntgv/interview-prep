package main

import (
	"reflect"
	"testing"
)

func TestSliceToLinkedList(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "empty",
			args: args{
				arr: []int{},
			},
			want: nil,
		},
		{
			name: "one",
			args: args{
				arr: []int{1},
			},
			want: &ListNode{Val: 1},
		},
		{
			name: "two",
			args: args{
				arr: []int{1, 2},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToLinkedList(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "one",
			args: args{
				head: LinkedListFromArgs(1),
				n:    1,
			},
			want: LinkedListFromArgs(),
		},
		{
			name: "two",
			args: args{
				head: LinkedListFromArgs(1, 2),
				n:    1,
			},
			want: LinkedListFromArgs(1),
		},
		{
			name: "five",
			args: args{
				head: LinkedListFromArgs(1, 2, 3, 4, 5),
				n:    1,
			},
			want: LinkedListFromArgs(1, 2, 3, 4),
		},
		{
			name: "five",
			args: args{
				head: LinkedListFromArgs(1, 2, 3, 4, 5),
				n:    4,
			},
			want: LinkedListFromArgs(1, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getListLength(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "two",
			args: args{
				head: LinkedListFromArgs(1, 2),
			},
			want: 2,
		},
		{
			name: "ten",
			args: args{
				head: LinkedListFromArgs(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getListLength(tt.args.head); got != tt.want {
				t.Errorf("getListLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
