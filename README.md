== Config Parser


Example:

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
    
    	
