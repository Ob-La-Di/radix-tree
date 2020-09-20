package main

import (
	"testing"
)

func Test_areLettersEqual(t *testing.T) {
	type args struct {
		path  string
		path2 string
		i     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return false",
			args: args{
				path:  "b",
				path2: "c",
				i:     0,
			},
			want: false,
		},

		{
			name: "should return true",
			args: args{
				path:  "c",
				path2: "c",
				i:     0,
			},
			want: true,
		},

		{
			name: "should return true",
			args: args{
				path:  "c",
				path2: "c",
				i:     1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreLettersEqual(tt.args.path, tt.args.path2, tt.args.i); got != tt.want {
				t.Errorf("areLettersEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_findNode(t *testing.T) {
//	type args struct {
//		node *Node
//		path string
//	}
//	tests := []struct {
//		name  string
//		args  args
//		want  *Node
//		want1 bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := findNode(tt.args.node, tt.args.path)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("findNode() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("findNode() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
func Test_shortestString(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return shortest string",
			args: args{
				a: "a",
				b: "abcd",
			},
			want: "a",
		},
		{
			name: "should return second parameter when both strings are the same length",
			args: args{
				a: "dcba",
				b: "abcd",
			},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestString(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("shortestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
