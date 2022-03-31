# Complex Flag Unmarshalling Helper

This library provides a facility to unmarshal complex flags from both inline and on-disk values, in JSON and YAML formats.

It is meant to be used alongside Jesse van den Keiboom's [Flags library](https://github.com/jessevdk/go-flags) to simplify the unmarshalling of values into Golang structs and arrays.

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
    return complexflag.UnmarshalInto(value, c)
}

```

The library provides support also for those cases where the exact type of the input data is not perfectly known in advance, or it may vary depending on e.g. a `type` field.

```golang
type CustomFlagType2 struct {
    // ...
}

func (c *CustomFlagType2) UnmarshalFlag(value string) error {
    var err error
    data, err = complexflag.Unmarshal(value)
    // after this call, data may contain a map[string]interface{} 
    // or a []interface{}, depending on whether the input is a 
    // JSON/YAML object; you can hook your custom unmarshalling 
    // logic here
    if t, ok := data["type"].(string); ok {
        switch(t) {
        case "foo":
            // do whatever to your CustomFlagType2 pointer
        case "bar":
            // ...
        }
    }
    return err
}

```