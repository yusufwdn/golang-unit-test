package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// table test
// adalh metode testing dengan menggunakan data di dalam slice/array (berupa name, request, danexpected)
// lalu melooping data tersebut sehingga bisa disebut dengan dinamis.
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Jamal",
			request:  "Jamal",
			expected: "Hello Jamal",
		},
		{
			name:     "Komarudin",
			request:  "Komarudin",
			expected: "Hello Komarudin",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			fmt.Println(result)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		result := "Hello A"
		require.Equal(t, "Hello A", result, "Result not match!")
	})

	t.Run("B", func(t *testing.T) {
		result := "Hello B"
		require.Equal(t, "Hello B", result, "Result not match!")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestAssert(t *testing.T) {
	result := "Hello Johny"
	// ketika gagal, assert akan mengembalikan fungsi Fail(). jadi, program akan dilanjutkan (jika ada).
	assert.Equal(t, "Hello Johny", result, "Result not match!")
}

func TestRequire(t *testing.T) {
	result := "Hello Johny"
	// ketika gagal, require akan mengembalikan fungsi FailNow(). artinya proses unit test tidak akan dilanjutkan (distop).
	require.Equal(t, "Hello Johny", result, "Result not match!")
}

func TestSkip(t *testing.T) {
	// fungsi t.Skip() akan melewatkan kode pada fungsi jika di dalamnya terdapat t.Skip()
	// namun sistem akan mengembalikan PASS setelah berhasil skip
	if "A" != "B" {
		t.Skip("Hello. This unit test will be skipped")
	}

	if runtime.GOOS == "darwin" {
		t.Skip("This program can't run in Mac Os")
	}

	result := HelloWorld("Johny")
	require.Equal(t, "Hello Johny", result)
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Jacky")

	if result != "Hello Jacky" {
		// error
		// mengembalikan nilai gagal namun akan melanjutkan ke fungsi selanjutnya.
		// t.Fail()

		// error
		// mengembalikan fail dengan pesan yang diisi pada arguments.
		t.Error("Result must be 'Hello Jacky'")
	}

	fmt.Println("Test Success")
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Brow!!")

	if result != "Hello Brow!!" {
		// error
		// mengembalikan error, dan tidak menjalankan fungsi selanjutnya
		// eksekusi unit test tidak akan dilanjutkan lagi
		// t.FailNow()

		t.Fatal("Result must be 'Hello Brow!!'")
	}

	fmt.Println("Test Success")
}
