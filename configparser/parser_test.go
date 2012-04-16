package configparser

import (
	"testing"
	"fmt"
)

func confEq(a ConfigFile, b ConfigFile) (err error) {
	for grpName, grpA := range a {
		grpB, found := b[grpName]
		if !found {
			return fmt.Errorf("Right side missing group %s", grpName)
		}

		for k, valA := range grpA {
			valB, found := grpB[k]

			if !found {
				return fmt.Errorf("Right side missing %s in [%s]", k, grpName)
			}

			if valA != valB {
				return fmt.Errorf("%s != %s in [%s].%s", valA, valB, k, grpName)
			}
		}
	}

	for grpName, grpA := range b {
		grpB, found := a[grpName]
		if !found {
			return fmt.Errorf("Left side missing group %s", grpName)
		}

		for k, _ := range grpB {
			_, found := grpA[k]

			if !found {
				return fmt.Errorf("Left side missing %s in [%s]", k, grpName)
			}
		}
	}

	return
}

var target1 = ConfigFile{
	"Cheese":{
		"foo":"bar",
		"baz":"d\ng\na",
	},
	"Burger":{
		"foo":"basg",
	},
	"Bummer:Man":{
		"this":"works",
	},
}


func TestParsing1(t *testing.T) {
	conf, err := ParseString(`
[Cheese]
foo = bar
baz = 
	d
	g
	a
[Burger]
foo = basg

[Bummer:Man]
this = works

}`,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("conf =", conf)

	err = confEq(conf, target1)
	if err != nil {
		t.Fatal(err)
	}

	t.Parallel()
}

func TestParsingColon(t *testing.T) {
	conf, err := ParseString(`
[Cheese]
foo: bar
baz : 
	d
	g
	a
[Burger]
foo = basg

[Bummer:Man]
this:
	works

}`,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("conf =", conf)

	err = confEq(conf, target1)
	if err != nil {
		t.Fatal(err)
	}

	t.Parallel()
}

func TestFile(t *testing.T) {
	conf, err := ParseFile("test.ini")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("conf =", conf)

	err = confEq(conf, target1)
	if err != nil {
		t.Fatal(err)
	}

	t.Parallel()
}
