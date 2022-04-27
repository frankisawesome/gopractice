package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}

func TestRepeat(t *testing.T) {
	t.Run("Test repeat a single time", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}

	})

	t.Run("Test repeat 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}

	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
