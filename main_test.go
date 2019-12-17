package main

import (
	"reflect"
	"testing"
)

func Test_isAFibonacciNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "base case 0", args: args{num: 0}, want: true},
		{name: "base case 1", args: args{num: 1}, want: true},
		{name: "not base case", args: args{num: 5}, want: true},
		{name: "not a valid case", args: args{num: 6}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAFibonacciNumber(tt.args.num); got != tt.want {
				t.Errorf("isAFibonacciNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCountListTo(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "base case", args: args{num: 0}, want: make([]int, 0, 0)},
		{name: "count to two", args: args{num: 2}, want: []int{1, 2}[:]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCountListTo(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCountListTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{name: "empty string", args: args{s: ""}, want: map[string]int{}},
		{name: "one word", args: args{s: "oi"}, want: map[string]int{"oi": 1}},
		{name: "two words, no repetition", args: args{s: "oi mundo"}, want: map[string]int{"oi": 1, "mundo": 1}},
		{name: "two words, with repetition", args: args{s: "oi mundo oi"}, want: map[string]int{"oi": 2, "mundo": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordCount(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
