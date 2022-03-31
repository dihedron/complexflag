package complexflag

import "testing"

func TestUnmarshalStructFromJSONFile(t *testing.T) {
	input := "@struct.json"
	result, err := Unmarshal(input)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result, ok := result.(map[string]interface{}); !ok {
		t.Fatalf("invalid output type: %T", result)
	} else {

		for k, v := range map[string]interface{}{
			"name":    "John",
			"surname": "Doe",
			"age":     float64(23),
		} {
			if result[k] != v {
				t.Errorf("error unmarshalling from file: expected %v (type %T) for key %v, got %v (type %T)", v, v, k, result[k], result[k])
			}
		}
	}
}

func TestUnmarshalArrayFromJSONFile(t *testing.T) {
	input := "@array.json"
	result, err := Unmarshal(input)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result, ok := result.([]interface{}); !ok {
		t.Fatalf("invalid output type: %T", result)
	} else {

		for i, v := range []interface{}{"one", "two", "three"} {
			if result[i] != v {
				t.Errorf("error unmarshalling from file: expected %v (type %T) for index %d, got %v (type %T)", v, v, i, result[i], result[i])
			}
		}
	}
}

func TestUnmarshalStructFromYAMLFile(t *testing.T) {
	input := "@struct.yaml"
	result, err := Unmarshal(input)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result, ok := result.(map[string]interface{}); !ok {
		t.Fatalf("invalid output type: %T", result)
	} else {

		for k, v := range map[string]interface{}{
			"name":    "John",
			"surname": "Doe",
			"age":     23,
		} {
			if result[k] != v {
				t.Errorf("error unmarshalling from file: expected %v (type %T) for key %v, got %v (type %T)", v, v, k, result[k], result[k])
			}
		}
	}
}

func TestUnmarshalArrayFromYAMLFile(t *testing.T) {
	input := "@array.yaml"
	result, err := Unmarshal(input)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result, ok := result.([]interface{}); !ok {
		t.Fatalf("invalid output type: %T", result)
	} else {

		for i, v := range []interface{}{"one", "two", "three"} {
			if result[i] != v {
				t.Errorf("error unmarshalling from file: expected %v (type %T) for index %d, got %v (type %T)", v, v, i, result[i], result[i])
			}
		}
	}
}

func TestUnmarshalStructFromJSONInline(t *testing.T) {
	input := `
	{
		"name": "John",
		"surname": "Doe",
		"age": 23
	}	
	`
	result, err := Unmarshal(input)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result, ok := result.(map[string]interface{}); !ok {
		t.Fatalf("invalid output type: %T", result)
	} else {

		for k, v := range map[string]interface{}{
			"name":    "John",
			"surname": "Doe",
			"age":     float64(23),
		} {
			if result[k] != v {
				t.Errorf("error unmarshalling from file: expected %v (type %T) for key %v, got %v (type %T)", v, v, k, result[k], result[k])
			}
		}
	}
}

func TestUnmarshalArrayFromJSONInline(t *testing.T) {
	input := `
	[
		"one",
		"two",
		"three"
	]
	`
	result, err := Unmarshal(input)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result, ok := result.([]interface{}); !ok {
		t.Fatalf("invalid output type: %T", result)
	} else {

		for i, v := range []interface{}{"one", "two", "three"} {
			if result[i] != v {
				t.Errorf("error unmarshalling from file: expected %v (type %T) for index %d, got %v (type %T)", v, v, i, result[i], result[i])
			}
		}
	}
}

func TestUnmarshalStructFromYAMLInline(t *testing.T) {
	input := `
---
name: John
surname: Doe
age: 23	
	`
	result, err := Unmarshal(input)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result, ok := result.(map[string]interface{}); !ok {
		t.Fatalf("invalid output type: %T", result)
	} else {

		for k, v := range map[string]interface{}{
			"name":    "John",
			"surname": "Doe",
			"age":     23,
		} {
			if result[k] != v {
				t.Errorf("error unmarshalling from file: expected %v (type %T) for key %v, got %v (type %T)", v, v, k, result[k], result[k])
			}
		}
	}
}

func TestUnmarshalArrayFromYAMLInline(t *testing.T) {
	input := `
---
- one
- two
- three	
	`
	result, err := Unmarshal(input)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result, ok := result.([]interface{}); !ok {
		t.Fatalf("invalid output type: %T", result)
	} else {

		for i, v := range []interface{}{"one", "two", "three"} {
			if result[i] != v {
				t.Errorf("error unmarshalling from file: expected %v (type %T) for index %d, got %v (type %T)", v, v, i, result[i], result[i])
			}
		}
	}
}

////////////////////

type s struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func TestUnmarshalIntoStructFromJSONFile(t *testing.T) {
	input := "@struct.json"
	result := &s{}
	err := UnmarshalInto(input, result)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result.Name != "John" {
		t.Fatalf("invalid value for name: expected John, got %v", result.Name)
	}
	if result.Surname != "Doe" {
		t.Fatalf("invalid value for name: expected Doe, got %v", result.Surname)
	}
	if result.Age != 23 {
		t.Fatalf("invalid value for name: expected 123, got %v", result.Age)
	}
}

func TestUnmarshalIntoArrayFromJSONFile(t *testing.T) {
	input := "@array.json"
	result := []string{}
	err := UnmarshalInto(input, &result)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	for i, v := range []interface{}{"one", "two", "three"} {
		if result[i] != v {
			t.Errorf("error unmarshalling from file: expected %v (type %T) for index %d, got %v (type %T)", v, v, i, result[i], result[i])
		}
	}
}

func TestUnmarshalIntoStructFromYAMLFile(t *testing.T) {
	input := "@struct.yaml"
	result := &s{}
	err := UnmarshalInto(input, result)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result.Name != "John" {
		t.Fatalf("invalid value for name: expected John, got %v", result.Name)
	}
	if result.Surname != "Doe" {
		t.Fatalf("invalid value for name: expected Doe, got %v", result.Surname)
	}
	if result.Age != 23 {
		t.Fatalf("invalid value for name: expected 123, got %v", result.Age)
	}
}

func TestUnmarshalIntoArrayFromYAMLFile(t *testing.T) {
	input := "@array.yaml"
	result := []string{}
	err := UnmarshalInto(input, &result)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	for i, v := range []interface{}{"one", "two", "three"} {
		if result[i] != v {
			t.Errorf("error unmarshalling from file: expected %v (type %T) for index %d, got %v (type %T)", v, v, i, result[i], result[i])
		}
	}
}

func TestUnmarshalIntoStructFromJSONInline(t *testing.T) {
	input := `
	{
		"name": "John",
		"surname": "Doe",
		"age": 23
	}
	`
	result := &s{}
	err := UnmarshalInto(input, result)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result.Name != "John" {
		t.Fatalf("invalid value for name: expected John, got %v", result.Name)
	}
	if result.Surname != "Doe" {
		t.Fatalf("invalid value for name: expected Doe, got %v", result.Surname)
	}
	if result.Age != 23 {
		t.Fatalf("invalid value for name: expected 123, got %v", result.Age)
	}
}

func TestUnmarshalIntoArrayFromJSONInline(t *testing.T) {
	input := `
	[
		"one",
		"two",
		"three"
	]
	`
	result := []string{}
	err := UnmarshalInto(input, &result)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	for i, v := range []interface{}{"one", "two", "three"} {
		if result[i] != v {
			t.Errorf("error unmarshalling from file: expected %v (type %T) for index %d, got %v (type %T)", v, v, i, result[i], result[i])
		}
	}
}

func TestUnmarshalIntoStructFromYAMLInline(t *testing.T) {
	input := `
---
name: John
surname: Doe
age: 23
	`
	result := &s{}
	err := UnmarshalInto(input, result)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	if result.Name != "John" {
		t.Fatalf("invalid value for name: expected John, got %v", result.Name)
	}
	if result.Surname != "Doe" {
		t.Fatalf("invalid value for name: expected Doe, got %v", result.Surname)
	}
	if result.Age != 23 {
		t.Fatalf("invalid value for name: expected 123, got %v", result.Age)
	}
}

func TestUnmarshalIntoArrayFromYAMLInline(t *testing.T) {
	input := `
---
- one
- two
- three
	`
	result := []string{}
	err := UnmarshalInto(input, &result)
	if err != nil {
		t.Fatalf("error unmarshalling from file: %v", err)
	}
	for i, v := range []interface{}{"one", "two", "three"} {
		if result[i] != v {
			t.Errorf("error unmarshalling from file: expected %v (type %T) for index %d, got %v (type %T)", v, v, i, result[i], result[i])
		}
	}
}
