package string

import "testing"

func TestHello(t *testing.T) {

	TEST := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("Say Hello To People", func(t *testing.T) {
		got := Hello("Viktor", "")
		want := "Hello, Viktor"
		TEST(t, got, want)
	})

	t.Run("Empty string 1", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		TEST(t, got, want)
	})

	t.Run("Empty string 2", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"
		TEST(t, got, want)
	})

	t.Run("Different Language1", func(t *testing.T) {
		got := Hello("Viktor", "Chinese")
		want := "哈囉，Viktor"
		TEST(t, got, want)
	})

	t.Run("Different Language2", func(t *testing.T) {
		got := Hello("Viktor", "English")
		want := "Hello, Viktor"
		TEST(t, got, want)
	})
}
