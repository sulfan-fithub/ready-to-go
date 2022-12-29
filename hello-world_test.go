package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Aligman", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Aligman")
		}
	})
	b.Run("Blink", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Blink")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Nap")
	}
}

func BenchmarkHelloWorldNapx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Napx")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(nap)",
			request:  "Nap",
			expected: "Hello Nap",
		},
		{
			name:     "HelloWorld(Sulfan)",
			request:  "Sulfan",
			expected: "Hello Sulfan",
		},
		{
			name:     "HelloWorld(Cali)",
			request:  "Cali",
			expected: "Hello Cali",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Nap", func(t *testing.T) {
		result := HelloWorld("Nap")
		require.Equal(t, "Hello Nap", result, "Result Must Be : Hello Nap")
	})
	t.Run("Sulfan", func(t *testing.T) {
		result := HelloWorld("Sulfan")
		require.Equal(t, "Hello Sulfan", result, "Result Must Be : Hello Nap")
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before Unit Test")
	m.Run()
	//after
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on Mac OS")
	}
	result := HelloWorld("aligman")
	require.Equal(t, "Hello aligman", result, "Result Must Be : Hello aligman")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("aligman")
	require.Equal(t, "Hello aligman", result, "Result Must Be : Hello aligman")
	fmt.Println("Test Hello World with Assert Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("aligman")
	assert.Equal(t, "Hello aligman", result, "Result Must Be : Hello aligman")
	fmt.Println("Test Hello World with Assert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("nap")
	if result != "Hello nap" {
		//unit test failed
		t.Fatal("Result must be : 'Hello nap'")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldSulfan(t *testing.T) {
	result := HelloWorld("Sulfan")
	if result != "Hello Sulfan" {
		//unit test failed
		t.Fatal("Result must be : 'Hello Sulfan'")
	}
	fmt.Println("TestHelloWorldSulfan Done")
}
