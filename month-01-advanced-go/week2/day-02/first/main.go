package main

import "fmt"

type Reader interface {
	Read() error
}

type Writer interface {
	Write() error
}

type Closer interface {
	Close() error
}

type Composition interface {
	Reader
	Writer
	Closer
}

type File struct {
	Name string
}

func (f File) Read() error {
	fmt.Println("Reading:", f.Name)
	return nil
}

func (f File) Write() error {
	fmt.Println("Writing:", f.Name)
	return nil
}

func (f File) Close() error {
	fmt.Println("Closing:", f.Name)
	return nil
}

func ProcessFile(f Composition) {
	_ = f.Read()
	_ = f.Write()
	_ = f.Close()
}

func main() {
	file := File{Name: "data.txt"}

	ProcessFile(file)
}
