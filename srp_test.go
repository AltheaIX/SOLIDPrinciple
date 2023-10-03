package main

import (
	"testing"
)

/*
File Management with SRP Principle
Checked and reviewed by ChatGPT
*/

func TestCreateFile(t *testing.T) {
	file, err := CreateFile("test.txt")
	if err != nil {
		t.Errorf("[TestCreateFile] %v - ERROR", err)
	}

	file.Close()
}

func TestReadFile(t *testing.T) {
	fileByte, err := ReadFile("test.txt")
	if err != nil {
		t.Errorf("[TestReadFile] %v - ERROR", err)
	}

	t.Logf("%v", string(fileByte))
}

func TestWriteToFile(t *testing.T) {
	multipleLine := `test
	multiple
	line`
	err := WriteToFile("test.txt", []byte(multipleLine))
	if err != nil {
		t.Errorf("[TestWriteToFile] %v - ERROR", err)
	}

	t.Log("[TestWriteToFile] Success")
}
