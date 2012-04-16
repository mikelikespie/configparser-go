ConfigParser for Go
===================

Implementation of a parser of Python's[ConfigParser] format for Go

[ConfigParser]: http://docs.python.org/library/configparser.html

Example usage:

    package main
    
    import (
    	"fmt"
    	"github.com/mikelikespie/configparser-go/configparser"
    )
    
    func main() {
    	// configparser.Parse and configparser.ParseFile also work
    	config, _ := configparser.ParseString(`[Hello]
    foo = bar
    boo : baz`)
    
    	foo := config["Hello"]["foo"]
    
    	fmt.Println(foo)
    }
    
    	
