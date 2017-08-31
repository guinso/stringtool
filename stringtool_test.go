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

func TestMakeSHA256(t *testing.T) {
	input := "hello world\n"
	expectedOutput := "a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447"

	output := MakeSHA256(input)

	if strings.Compare(output, expectedOutput) != 0 {
		t.Errorf("Expected %s, got %s instead. input = %s", expectedOutput, output, input)
	}
}
