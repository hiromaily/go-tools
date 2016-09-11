package {{.Name}}_test

import (
	. "github.com/hiromaily/golibs/{{.Name}}"
	//lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"testing"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[{{.Uppercase}}]")
}

func setup() {
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// function
//-----------------------------------------------------------------------------


//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func Test{{.Uppercase}}(t *testing.T) {
	//if err != nil {
	//	t.Errorf("Test{{.Uppercase}} error: %s", err)
	//}
}


//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func Benchmark{{.Uppercase}}(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}

