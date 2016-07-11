package {{.Name}}_test

import (
	"flag"
	. "github.com/hiromaily/golibs/{{.Name}}"
	lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[{{.Uppercase}}_TEST]", "/var/log/go/test.log")
	if *benchFlg == 0 {
	}
}

func teardown() {
	if *benchFlg == 0 {
	}
}

// Initialize
func TestMain(m *testing.M) {
	flag.Parse()

	//TODO: According to argument, it switch to user or not.
	//TODO: For bench or not bench
	setup()

	code := m.Run()

	teardown()

	// 終了
	os.Exit(code)
}

//-----------------------------------------------------------------------------
// {{.Uppercase}}
//-----------------------------------------------------------------------------
func Test{{.Uppercase}}(t *testing.T) {
	if *benchFlg == 1 {
		t.Skip("skipping Test{{.Uppercase}}")
	}

	//if err != nil {
	//	t.Errorf("Test{{.Uppercase}} error: %s", err)
	//}

}

//-----------------------------------------------------------------------------
//Benchmark
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

