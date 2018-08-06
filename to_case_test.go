package namecase

import (
	"testing"
)

var testdataBascis = []string{
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
	"hello world",

	"Hello world",
}

func TestToUpperSnake(t *testing.T) {
	for _, v := range testdataBascis {
		d := ToUpperSnake(v)
		if d != "HELLO_WORLD" {
			t.Error(v, d)
		}
	}
}

func TestToLowerSnake(t *testing.T) {
	for _, v := range testdataBascis {
		d := ToLowerSnake(v)
		if d != "hello_world" {
			t.Error(v, d)
		}
	}
}

func TestToUpperStrike(t *testing.T) {
	for _, v := range testdataBascis {
		d := ToUpperStrike(v)
		if d != "HELLO-WORLD" {
			t.Error(v, d)
		}
	}
}

func TestToLowerStrike(t *testing.T) {
	for _, v := range testdataBascis {
		d := ToLowerStrike(v)
		if d != "hello-world" {
			t.Error(v, d)
		}
	}
}

func TestToUpperSpace(t *testing.T) {
	for _, v := range testdataBascis {
		d := ToUpperSpace(v)
		if d != "HELLO WORLD" {
			t.Error(v, d)
		}
	}
}

func TestToLowerSpace(t *testing.T) {
	for _, v := range testdataBascis {
		d := ToLowerSpace(v)
		if d != "hello world" {
			t.Error(v, d)
		}
	}
}

func TestToUpperHump(t *testing.T) {
	for _, v := range testdataBascis {
		d := ToUpperHump(v)
		if d != "HelloWorld" {
			t.Error(v, d)
		}
	}
}

func TestToLowerHump(t *testing.T) {
	for _, v := range testdataBascis {
		d := ToLowerHump(v)
		if d != "helloWorld" {
			t.Error(v, d)
		}
	}
}

var testdataInitialisms = []string{
	"user id",
	"userID",
	"userId",
	"user_id",
}

func TestLowerInitialisms(t *testing.T) {
	for _, v := range testdataInitialisms {
		d := ToLowerHumpInitialisms(v)
		if d != "userID" {
			t.Error(v, d)
		}
	}
}

func TestUpperInitialisms(t *testing.T) {
	for _, v := range testdataInitialisms {
		d := ToUpperHumpInitialisms(v)
		if d != "UserID" {
			t.Error(v, d)
		}
	}
}
