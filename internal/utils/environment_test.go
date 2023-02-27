package utils

import (
	"os"
	"testing"
)

func TestGetEnvVariableIsSet(t *testing.T) {
	variableName := "TEST_ENV_VAR"
	defer os.Unsetenv(variableName)
	os.Setenv(variableName, "someSetEnvVariable")

	myVariable := GetEnv(variableName, "defaultEnvVariable")
	expected := "someSetEnvVariable"
	if myVariable != expected {
		t.Errorf("Function returned wrong result: got %v wanted %v",
			myVariable, expected)
	}
}

func TestGetEnvVariableIsNotSet(t *testing.T) {
	variableName := "TEST_ENV_VAR"
	os.Unsetenv(variableName)

	myVariable := GetEnv(variableName, "defaultEnvVariable")
	expected := "defaultEnvVariable"
	if myVariable != expected {
		t.Errorf("Function returned wrong result: got %v wanted %v",
			myVariable, expected)
	}
}
