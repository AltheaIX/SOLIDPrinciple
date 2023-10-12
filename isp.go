package main

import "fmt"

/*
Document Management System with Single Responsibility Principle, Open-Closed Principle, Liskov Substitution Principle and Interface Segregation Principle
Checked and Reviewed by ChatGPT
*/

type DocumentReader interface {
	Open() string
}

type DocumentCloser interface {
	Close() string
}

type DocumentSaver interface {
	Save() string
}

type TextDocument struct {
}

func (td *TextDocument) Open() string {
	return "[TextDocumentOpen] Document has been opened."
}

func (td *TextDocument) Close() string {
	return "[TextDocumentClose] Document has been closed."
}

func (td *TextDocument) Save() string {
	return "[TextDocumentSave] Document has been saved."
}

func NewTextDocumentReader() DocumentReader {
	return &TextDocument{}
}

func NewTextDocumentCloser() DocumentCloser {
	return &TextDocument{}
}

func NewTextDocumentSaver() DocumentSaver {
	return &TextDocument{}
}

type SpreadsheetDocument struct {
}

func (sd *SpreadsheetDocument) Open() string {
	return "[SpreadsheetDocumentOpen] Document has been opened."
}

func (sd *SpreadsheetDocument) Close() string {
	return "[SpreadsheetDocumentClose] Document has been closed."
}

func (sd *SpreadsheetDocument) Save() string {
	return "[SpreadsheetDocumentSave] Document has been saved."
}

func NewSpreadsheetDocumentReader() DocumentReader {
	return &SpreadsheetDocument{}
}

func NewSpreadsheetDocumentCloser() DocumentCloser {
	return &SpreadsheetDocument{}
}

func NewSpreadsheetDocumentSaver() DocumentSaver {
	return &SpreadsheetDocument{}
}

type Printable interface {
	Print()
}

type PrintManager struct {
	reader DocumentReader
}

func (pm *PrintManager) Print() {
	fmt.Println("Inside of the document:")
	fmt.Println(pm.reader.Open())
}

func NewPrintManager(reader DocumentReader) Printable {
	return &PrintManager{reader}
}

type Signable interface {
	AddSignature() string
	VerifySignature() string
}

type SignManager struct {
	reader DocumentReader
	saver  DocumentSaver
	closer DocumentCloser
}

func (sm *SignManager) AddSignature() string {
	sm.reader.Open()
	sm.saver.Save()
	sm.closer.Close()
	return "[SignManagerAddSignature] Signature has been added to the document."
}

func (sm *SignManager) VerifySignature() string {
	sm.reader.Open()
	fmt.Println("[VerifySignature] Verifying, please to wait.")
	sm.closer.Close()
	return "[VerifySignature] Success verifying the signature of document."
}

func NewSignManager(reader DocumentReader, saver DocumentSaver, closer DocumentCloser) Signable {
	return &SignManager{reader: reader, saver: saver, closer: closer}
}

func ProcessDocument(reader DocumentReader, saver DocumentSaver, closer DocumentCloser, printManager Printable, signManager Signable) {
	fmt.Println(reader.Open())
	fmt.Println(saver.Save())
	fmt.Println(closer.Close())
	printManager.Print()
	fmt.Println(signManager.AddSignature())
}
