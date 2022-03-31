# Flexible unmarshalling of value sComplex Flag Unmarshalling Helper

[![Go Report Card](https://goreportcard.com/badge/github.com/dihedron/unknown)](https://goreportcard.com/report/github.com/dihedron/unknown)
[![Godoc](https://godoc.org/github.com/dihedron/unknown?status.svg)](https://godoc.org/github.com/dihedron/unknown)

This library provides a facility to unmarshal unknown values from both inline and on-disk values, in either JSON or YAML formats.

It can be used wherever an input must be unmarshalled into a Golang object.

## Using with https://github.com/jessevdk/go-flags

One such use is alongside Jesse van den Keiboom's [Flags library](https://github.com/jessevdk/go-flags), to simplify the unmarshalling of values into Golang structs and arrays.

```golang
type MyCommand struct {
    Param1     CustomFlagType1   `short:"p" long:"param1" description:"An input parameter, either as an inline value or as a @file (in JSON or YAML format)."`
    Param2     CustomFlagType2   `short:"q" long:"param2" description:"A partially deserialised input parameter, either as an inline value or as a @file (in JSON or YAML format)."`
}

type CustomFlagType1 struct {
    Name    string `json:"name,omitempty" yaml:"name,omitempty"`
    Surname string `json:"surname,omitempty" yaml:"surname,omitempty"`
    Age     int    `json:"age,omitempty" yaml:"age,omitempty"`
}

func (c *CustomFlagType1) UnmarshalFlag(value string) error {
    return unknown.UnmarshalInto(value, c)
}

```

The library provides support also for those cases where the exact type of the input data is not perfectly known in advance, or it may vary depending on e.g. a `type` field.

Note that the `UnmarshalInto` function expects a pointer to the destination struct/array to be passed in, so the object and the input value must both be known in advance and match one another. 
The `Unmarshal` function is more lax: it detects the type of entity (object/array) in the input and *returns* either a `map[string]interface{}` (if the input value is an object) or a `[]interface{}` if the input is an array of objects. It is up to the caller to handle the two cases properly, but it leaves the possibility of using e.g. such tools as Mitchell Hashimoto's [Map Structure](https://github.com/mitchellh/mapstructure) library to perform smarter, adaptive staged unmarshalling. 

```golang
type CustomFlagType2 struct {
    // ...
}

func (c *CustomFlagType2) UnmarshalFlag(value string) error {
    var err error
    data, err = complexflag.Unmarshal(value)
    // after this call, data may contain a map[string]interface{} 
    // or a []interface{}, depending on whether the input is a 
    // JSON/YAML object or an array; you can hook your custom 
    // unmarshalling logic here
    switch data := data.(type) {
    case map[string]interface{}:
        // it's a map, switch on the "type" field
        // retrieve the "type" value and cast it to string, 
        if v, ok := data["type"] ; ok {
            // the value is there, attempt casting it to string
            if t, ok := v.(string); ok {
                switch(t) {
                case "foo":
                    // do whatever you need to do with a type "foo";
                    // the CustomFlagType2 pointer allows manipulation 
                    // of the struct
                    c.SomeField = data["key_dependent_on_type_foo"]
                    // and so on...
                case "bar":
                    // ...same for "bar"
                }
            }
        }
    case []interface{}:
	default:
		return errors.New("unexpected type of returned data")
	}    
    return err
}

```

## Importing the library

In order to use the library, import it like this:

```golang
import (
    "github.com/dihedron/unknown"
)
```

Then open a command prompt in your project's root directory and run:

```bash
$> go mod tidy
```

