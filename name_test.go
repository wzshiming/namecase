package nomenclature

import "testing"

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

	// TODO:
	//	"HELLO_WORLD",
	//	"_HELLO_WORLD_",
	//	"_HELLO____WORLD_",
	//	"_HEllo__WORLD_",
}

var data2 = []string{
	"a_id",
	"AId",
	"A_Id",
}

func TestName(t *testing.T) {
	for _, v := range data1 {
		d := Hump2Snake(v)
		if d != "hello_world" {
			t.Error(v, d)
		}
	}
	for _, v := range data1 {
		d := Snake2Hump(v)
		if d != "HelloWorld" {
			t.Error(v, d)
		}
	}

	for _, v := range data2 {
		d := Snake2Hump(v)
		if d != "AId" {
			t.Error(v, d)
		}
	}

	for _, v := range data2 {
		d := Hump2Snake(v)
		if d != "a_id" {
			t.Error(v, d)
		}
	}
}
