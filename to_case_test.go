package namecase

import (
	"testing"
)

var data1 = []string{
	"helloWorld",
	"HelloWorld",
	"hello_world",

	"_helloWorld_",
	"_HelloWorld_",
	"_hello_world_",

	"_hello____world_",
	"_hello__World_",
	"_Hello__World_",

	"HELLO_WORLD",
	"_HELLO_WORLD_",
	"_HELLO____WORLD_",
	"_HEllo__WORLD_",

	"hello world",

	"Hello world",
}

func TestToUpperSnake(t *testing.T) {
	for _, v := range data1 {
		d := ToUpperSnake(v)
		if d != "HELLO_WORLD" {
			t.Error(v, d)
		}
	}
}

func TestToLowerSnake(t *testing.T) {
	for _, v := range data1 {
		d := ToLowerSnake(v)
		if d != "hello_world" {
			t.Error(v, d)
		}
	}
}

func TestToUpperStrike(t *testing.T) {
	for _, v := range data1 {
		d := ToUpperStrike(v)
		if d != "HELLO-WORLD" {
			t.Error(v, d)
		}
	}
}

func TestToLowerStrike(t *testing.T) {
	for _, v := range data1 {
		d := ToLowerStrike(v)
		if d != "hello-world" {
			t.Error(v, d)
		}
	}
}

func TestToUpperSpace(t *testing.T) {
	for _, v := range data1 {
		d := ToUpperSpace(v)
		if d != "HELLO WORLD" {
			t.Error(v, d)
		}
	}
}

func TestToLowerSpace(t *testing.T) {
	for _, v := range data1 {
		d := ToLowerSpace(v)
		if d != "hello world" {
			t.Error(v, d)
		}
	}
}

func TestToUpperHump(t *testing.T) {
	for _, v := range data1 {
		d := ToUpperHump(v)
		if d != "HelloWorld" {
			t.Error(v, d)
		}
	}
}

func TestToLowerHump(t *testing.T) {
	for _, v := range data1 {
		d := ToLowerHump(v)
		if d != "helloWorld" {
			t.Error(v, d)
		}
	}

}
