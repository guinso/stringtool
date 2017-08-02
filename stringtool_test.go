package stringtool

import (
	"strings"
	"testing"
)

func TestSnakeToCamelCase(t *testing.T) {
	input := "snake_case"
	expectedOutput := "SnakeCase"

	output := SnakeToCamelCase(input)

	if strings.Compare(output, expectedOutput) != 0 {
		t.Errorf("Expect %s, got %s intead. input = %s",
			expectedOutput, output, input)
	}
}

func TestToSnakeCase(t *testing.T) {
	input := "camelCase"
	expectedOutput := "camel_case"

	output := ToSnakeCase(input)

	if strings.Compare(output, expectedOutput) != 0 {
		t.Errorf("Expect %s, got %s intead. input = %s",
			expectedOutput, output, input)
	}
}
