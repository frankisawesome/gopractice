package generics_prac

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual[int](t, 1, 1)
		AssertNotEqual[int](t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual[string](t, "hello", "hello")
		AssertNotEqual[string](t, "hello", "grace")
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %d", got)
	}
}
