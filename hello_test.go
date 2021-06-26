package main

import "testing"

// Test function name should be start with "Test" keyword
// It only takes one arguments *testing.T
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function call
		// rather than inside our test helper. This will help other developers track down
		// problems easier.
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// Here we are introducing another tool in our testing arsenal,
	//subtests. Sometimes it is useful to group tests around a "thing"
	//and then have subtests describing different scenarios.
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		if got != want {
			// t.Errorf() print error to stdout and fail the test.
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		if got != want {
			assertCorrectMessage(t, got, want)
		}

	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
