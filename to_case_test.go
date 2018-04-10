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
}

func TestToLowerSnake(t *testing.T) {
	for _, v := range data1 {
		d := ToLowerSnake(v)
		if d != "hello_world" {
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

func TestToUpperSnake(t *testing.T) {
	for _, v := range data1 {
		d := ToUpperSnake(v)
		if d != "HELLO_WORLD" {
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
