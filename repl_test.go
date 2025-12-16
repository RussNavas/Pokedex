package main

import (
	"testing"
)

func TestCleanInput(t *testing.T){
	/*
		split the user's input into "words" based on whitespace. It should also
		lowercase the input and trim any leading or trailing whitespace.
		For example:
			hello world -> ["hello", "world"]
			Charmander Bulbasaur PIKACHU -> ["charmander", "bulbasaur", "pikachu"]
	*/

	cases := []struct {
		input		string
		expected	[]string
	}{
		{
			input:		" hello world ",
			expected: 	[]string{"hello", "world"},
		},
		{
			input:		"Charmander Bulbasaur PIKACHU",
			expected:	[]string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:		"  cHarManDer       bulbasaur PikaChu",
			expected:	[]string{"charmander", "bulbasaur", "pikachu"},
		},
	}


	// Loop over cases and run the tests:
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected){
			t.Fatalf("expected: %v\n actual: %v\n", c.expected, actual )
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord{
				t.Fatalf("expected: %v\n actual: %v\n", word, expectedWord)
			}
		}
	}
}

func TestCommands(t *testing.T) {
    // 1. Get the map of commands (you might need to export this or make a getter)
    commands := getCommands() 

    // 2. Define the cases you want to check exist
    expected := []string{"help", "exit"}

    // 3. Loop through expected commands and ensure they exist in the map
    for _, name := range expected {
        cmd, ok := commands[name]
        if !ok {
            t.Errorf("Command %s not found in registry", name)
            continue
        }
        
        // Optional: Check if the name matches the key
        if cmd.name != name {
             t.Errorf("Expected command name %s, got %s", name, cmd.name)
        }
    }
}
