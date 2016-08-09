package {{.Name}}_test

import (
	. "github.com/hiromaily/golibs/{{.Name}}"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
)

var (
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[{{.Uppercase}}_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}
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

