package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myReaderBelt:= bufio.NewReader(os.Stdin)
	fmt.Print("Please enter q quote to save to file -> ")
	myquote , _:= myReaderBelt.ReadString('\n')
	addQuote(myquote)
	
}

func addQuote(myquote string){
	fileHandler,_:=os.OpenFile("quotes.txt",os.O_RDWR|os.O_APPEND|os.O_CREATE,0666)

	// after exit of main close the file
	defer fileHandler.Close()

	// we need a belt to convey buffered data
	// for fileHandler
	myWriterBelt := bufio.NewWriter(fileHandler)

	// our intention is through the belt to 
	// move our data (myqote) to file
	fmt.Fprintln(myWriterBelt,myquote)
	
	// write dashes
	fmt.Fprintln(myWriterBelt,"-------------")

	// move the data to file
	myWriterBelt.Flush()

	fmt.Println("Data has been saved to quotes.txt")
}
