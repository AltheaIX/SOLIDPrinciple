package main

import "os"

/*
File Management with SRP Principle
Checked and reviewed by ChatGPT
*/

func CreateFile(name string) (*os.File, error) {
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ReadFile(name string) ([]byte, error) {
	fileByte, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return fileByte, nil
}

func WriteToFile(name string, toWrite []byte) error {
	err := os.WriteFile(name, toWrite, 0666)
	if err != nil {
		return err
	}

	return nil
}
