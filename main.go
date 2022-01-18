package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myquote := `Amateurs practice until they get it right.
Proffesionals practice until they cant't get it wrong.`
	// - write quotes.txt to hard disk
	// - return a file handler for the file
	fileHandler,_:=os.Create("quotes.txt")

	// after exit of main close the file
	defer fileHandler.Close()

	// we need a belt to convey buffered data
	myWriterBelt := bufio.NewWriter(fileHandler)

	// our intention is through the belt to 
	// move our data (myqote) to file
	fmt.Fprintln(myWriterBelt,myquote)

	// move the data to file
	myWriterBelt.Flush()

	fmt.Println("Data has been saved to quotes.txt")
}
