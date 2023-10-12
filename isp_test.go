package main

import "testing"

/*
Document Management System with Single Responsibility Principle, Open-Closed Principle, Liskov Substitution Principle and Interface Segregation Principle
Checked and Reviewed by ChatGPT
*/

func TestTextDocument_Open(t *testing.T) {
	td := NewTextDocumentReader()
	t.Log(td.Open())
}

func TestTextDocument_Close(t *testing.T) {
	td := NewTextDocumentCloser()
	t.Log(td.Close())
}

func TestTextDocument_Save(t *testing.T) {
	td := NewTextDocumentSaver()
	t.Log(td.Save())
}

func TestSpreadsheetDocument_Open(t *testing.T) {
	sd := NewSpreadsheetDocumentReader()
	t.Log(sd.Open())
}

func TestSpreadsheetDocument_Close(t *testing.T) {
	sd := NewSpreadsheetDocumentCloser()
	t.Log(sd.Close())
}

func TestSpreadsheetDocument_Save(t *testing.T) {
	sd := NewSpreadsheetDocumentSaver()
	t.Log(sd.Save())
}

func TestPrintManagerTextDocument_Print(t *testing.T) {
	pm := NewPrintManager(&TextDocument{})
	pm.Print()
}

func TestPrintManagerSpreadsheetDocument_Print(t *testing.T) {
	pm := NewPrintManager(&SpreadsheetDocument{})
	pm.Print()
}

func TestSignManagerTextDocument_AddSignature(t *testing.T) {
	reader := NewTextDocumentReader()
	saver := NewTextDocumentSaver()
	closer := NewTextDocumentCloser()
	sm := NewSignManager(reader, saver, closer)

	t.Log(sm.AddSignature())
}

func TestSignManagerTextDocument_VerifySignature(t *testing.T) {
	reader := NewTextDocumentReader()
	saver := NewTextDocumentSaver()
	closer := NewTextDocumentCloser()
	sm := NewSignManager(reader, saver, closer)

	t.Log(sm.VerifySignature())
}

func TestSignManagerSpreadsheetDocument_AddSignature(t *testing.T) {
	reader := NewSpreadsheetDocumentReader()
	saver := NewSpreadsheetDocumentSaver()
	closer := NewSpreadsheetDocumentCloser()
	sm := NewSignManager(reader, saver, closer)

	t.Log(sm.AddSignature())
}

func TestSignManagerSpreadsheetDocument_VerifySignature(t *testing.T) {
	reader := NewSpreadsheetDocumentReader()
	saver := NewSpreadsheetDocumentSaver()
	closer := NewSpreadsheetDocumentCloser()
	sm := NewSignManager(reader, saver, closer)

	t.Log(sm.VerifySignature())
}

func TestProcessDocumentTextDocument(t *testing.T) {
	reader := NewTextDocumentReader()
	saver := NewTextDocumentSaver()
	closer := NewTextDocumentCloser()
	pm := NewPrintManager(reader)
	sm := NewSignManager(reader, saver, closer)

	ProcessDocument(reader, saver, closer, pm, sm)
}

func TestProcessDocumentSpreadsheetDocument(t *testing.T) {
	reader := NewSpreadsheetDocumentReader()
	saver := NewSpreadsheetDocumentSaver()
	closer := NewSpreadsheetDocumentCloser()
	pm := NewPrintManager(reader)
	sm := NewSignManager(reader, saver, closer)

	ProcessDocument(reader, saver, closer, pm, sm)
}
