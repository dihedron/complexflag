package complexflag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

// Format is the type representing the possible formats
// for the complex flag structure.
type Format uint8

const (
	// FormatJSON indicates that the flag is in JSON format.
	FormatJSON Format = iota
	// FormatYAML indicates that the flag is in YAML format.
	FormatYAML
)

// Unmarshal unmarshals a complex command line flag into an argument;
// if the command line argument starts with a '@' it is assumed to
// be a file on the local filesystem, it is read into memory and then
// unmarshalled into the object struct, which must be appropriately
// annotated; if it does not start with '@', it can be either a YAML
// inline representation (in which case it MUST start with '---') or
// and inline JSON representation and is unmarshalled acoordingly.
func Unmarshal(value string) (interface{}, error) {
	format := FormatJSON
	var content []byte
	if strings.HasPrefix(value, "@") {
		// it's a file on disk, check it exist
		filename := strings.TrimPrefix(value, "@")
		info, err := os.Stat(filename)
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file '%s' does not exist: %w", filename, err)
		}
		if info.IsDir() {
			return nil, fmt.Errorf("'%s' is a directory, not a file", filename)
		}
		// read into memory
		content, err = ioutil.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("error reading file '%s': %w", filename, err)
		}
		// type detection is based on file extension
		ext := path.Ext(filename)
		switch strings.ToLower(ext) {
		case ".yaml", ".yml":
			format = FormatYAML
		case ".json":
			format = FormatJSON
		default:
			return nil, fmt.Errorf("unsupported data format in file: %s", path.Ext(filename))
		}
	} else {
		// not a file, type detection is based on the data
		value = strings.TrimSpace(value)
		content = []byte(value)
		if strings.HasPrefix(value, "---") {
			format = FormatYAML
		} else if strings.HasPrefix(value, "{") || strings.HasPrefix(value, "[") {
			// TODO: we could optimise by recording whether it's a struct or an array
			format = FormatJSON
		} else {
			return nil, fmt.Errorf("unrecognisable input format in inline data")
		}
	}
	// now depending on the format, unmarshal to JSON or YAML
	switch format {
	case FormatJSON:
		// NOTE: a JSON document can represent either an object or an array
		// bot the standard library methods expect the target object to be
		// pre-allocated; thus, we try to unmarshal to a map, which is the
		// most general representation of a struct; if it fails with a parse
		// error because the JSON document represents an array, we try with
		// an array next

		// first attempt: unmarshalling to a map (like a struct would)...
		m := map[string]interface{}{}
		if err := json.Unmarshal(content, &m); err != nil {
			if err, ok := err.(*json.UnmarshalTypeError); ok {
				if err.Value == "array" && err.Offset == 1 {
					// second attempt: it is not a struct, it's an array, let's try that...
					a := []interface{}{}
					if err := yaml.Unmarshal(content, &a); err != nil {
						return nil, fmt.Errorf("error unmarshalling from JSON: %w", err)
					}
					return a, nil
				}
			}
			return nil, fmt.Errorf("error unmarshalling from JSON: %w", err)
		}
		return m, nil
	case FormatYAML:
		object := map[string]interface{}{}
		if err := yaml.Unmarshal(content, object); err != nil {
			if err, ok := err.(*yaml.TypeError); ok {
				// TODO: find a way to circumvent marshalling error in case of array
				for _, e := range err.Errors {
					if strings.HasSuffix(e, "cannot unmarshal !!seq into map[string]interface {}") {
						// second attempt: it is not a struct, it's an array, let's try that...
						a := []interface{}{}
						if err := yaml.Unmarshal(content, &a); err != nil {
							return nil, fmt.Errorf("error unmarshalling from YAML: %w", err)
						}
						return a, nil
					}
				}
				return nil, fmt.Errorf("error: %s, %+v", err.Error(), err.Errors)
			}
			return nil, fmt.Errorf("error unmarshalling from YAML: %w (%T)", err, err)

		}
		return object, nil
	default:
		return nil, fmt.Errorf("unsupported encoding: %v", format)
	}
}

func UnmarshalInto(value string, target interface{}) error {
	format := FormatJSON
	var content []byte
	if strings.HasPrefix(value, "@") {
		// it's a file on disk, check it exist
		filename := strings.TrimPrefix(value, "@")
		info, err := os.Stat(filename)
		if os.IsNotExist(err) {
			return fmt.Errorf("file '%s' does not exist: %w", filename, err)
		}
		if info.IsDir() {
			return fmt.Errorf("'%s' is a directory, not a file", filename)
		}
		// read into memory
		content, err = ioutil.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("error reading file '%s': %w", filename, err)
		}
		// type detection is based on file extension
		ext := path.Ext(filename)
		switch strings.ToLower(ext) {
		case ".yaml", ".yml":
			format = FormatYAML
		case ".json":
			format = FormatJSON
		default:
			return fmt.Errorf("unsupported data format in file: %s", path.Ext(filename))
		}
	} else {
		// not a file, type detection is based on the data
		value = strings.TrimSpace(value)
		content = []byte(value)
		if strings.HasPrefix(value, "---") {
			format = FormatYAML
		} else if strings.HasPrefix(value, "{") || strings.HasPrefix(value, "[") {
			// TODO: we could optimise by recording whether it's a struct or an array
			format = FormatJSON
		} else {
			return fmt.Errorf("unrecognisable input format in inline data")
		}
	}
	// now depending on the format, unmarshal to JSON or YAML
	switch format {
	case FormatJSON:
		if err := json.Unmarshal(content, target); err != nil {
			return fmt.Errorf("error unmarshalling from JSON: %w", err)
		}
		return nil
	case FormatYAML:
		if err := yaml.Unmarshal(content, target); err != nil {
			return fmt.Errorf("error unmarshalling from YAML: %w (%T)", err, err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported encoding: %v", format)
	}
}
