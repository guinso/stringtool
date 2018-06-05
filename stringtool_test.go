package stringtool

import (
	"strings"
	"testing"

	"github.com/satori/go.uuid"
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

func TestGenerateRandomUUID(t *testing.T) {
	uuid1, err1 := uuid.NewV4()
	uuid2, err2 := uuid.NewV4()

	if err1 != nil {
		t.Error(err1)
		return
	}

	if err2 != nil {
		t.Error(err2)
		return
	}

	if strings.Compare(uuid1.String(), uuid2.String()) == 0 {
		t.Errorf("both UUID value cannot be identical")
	}
}
