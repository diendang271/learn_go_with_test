package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertionRepeatMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	}
	t.Run("test repeated 5 times", func(t *testing.T) {
		got := Repeat("A", 5)
		want := "AAAAA"
		assertionRepeatMessage(t, got, want)
	})
	t.Run("test repeated 0 times", func(t *testing.T) {
		got := Repeat("A", 0)
		want := ""
		assertionRepeatMessage(t, got, want)
	})
}

func ExampleRepeat()  {
	result := Repeat("A", 5)
	fmt.Println(result)
	// Output: AAAAA
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}