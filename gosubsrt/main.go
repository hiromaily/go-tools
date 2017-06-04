package main

import (
	"bufio"
	"flag"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	tim "github.com/hiromaily/golibs/time"
	u "github.com/hiromaily/golibs/utils"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

var (
	file      = flag.String("f", "", "path of str file")
	tweakTime = flag.Float64("t", 0.0, "path of srt file")
	tmpPath   = "/tmp/tempfile"
	outPath   = "out"
)

var usage = `Usage: %s [options...]
Options:
  -f  path of srt file.
  -t  time of tweaking duration.
e.g.:
  gosubsrt -f ./xxxxx.srt -t -1.5
`

func readSrtFile() error {
	//1.Read
	fp, err := os.Open(*file)
	//fp, err := os.OpenFile(*file, os.O_RDWR, 0660)
	if err != nil {
		lg.Error(err)
		return err
	}
	defer fp.Close()

	//file size
	fi, err := fp.Stat()
	if err != nil {
		lg.Error(err)
		return err
	} else if fi.Size() > math.MaxInt32 {
		lg.Error(err)
		return fmt.Errorf("file size is too big: %d", fi.Size())
	}

	//fmt.Printf("file name: %s\n", fi.Name())
	//fmt.Printf("file size(byte): %d\n", fi.Size())
	//fmt.Printf("mode: %s\n", fi.Mode())
	//fmt.Printf("directory :%t\n", fi.IsDir())

	//different pattern to file
	//scanner := bufio.NewScanner(fp)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}

	reader := bufio.NewReaderSize(fp, int(fi.Size()))

	//2.Write
	fp2, err := os.Create(tmpPath)
	if err != nil {
		lg.Error(err)
		return err
	}
	defer fp2.Close()
	writer := bufio.NewWriter(fp2)

	for {
		line, _, res := reader.ReadLine()
		//fmt.Println(string(line))
		//
		reWrite(writer, string(line))

		if res == io.EOF {
			break
		} else if res != nil {
			lg.Error(err)
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		//invalid argument
		lg.Error(err)
		return err
	}

	return nil
}

func reWrite(writer *bufio.Writer, line string) {
	//modify
	//00:00:10,950 --> 00:00:14,490
	buf := strings.Split(line, " --> ")
	if len(buf) == 2 {
		//rewirte
		buf[0] = calcTime(buf[0])
		buf[1] = calcTime(buf[1])
		writer.WriteString(fmt.Sprintf("%s --> %s\n", buf[0], buf[1]))
	} else {
		writer.WriteString(line + "\n")
	}
}

func calcTime(timeStr string) string {
	//00:00:10,950
	tims := strings.Split(strings.Replace(timeStr, ",", ":", -1), ":")
	timI := u.ConvertToInt(tims)

	ti := tim.GetFormatTime2(timI[0], timI[1], timI[2], timI[3]*int(math.Pow10(6)))

	integerVal := math.Trunc(*tweakTime)
	decimalVal := math.Trunc((*tweakTime - math.Trunc(*tweakTime)) * 1000)
	if integerVal != 0 {
		ti = ti.Add(time.Duration(integerVal) * time.Second)
	}
	if decimalVal != 0 {
		ti = ti.Add(time.Duration(int(decimalVal)*int(math.Pow10(6))) * time.Nanosecond)
	}
	return strings.Replace(ti.Format("15:04:05.000"), ".", ",", -1)
}

func copyFile() error {
	//move
	outPath = strings.Replace(*file, "srtfiles", outPath, -1)
	//err := os.Rename(tmpPath, *file)
	err := os.Rename(tmpPath, outPath)
	if err != nil {
		//invalid argument
		lg.Error(err)
		return err
	}
	return nil
}

func init() {
	lg.InitializeLog(lg.DebugStatus, lg.LogOff, 99, "[GOTOOLS GoTestFile]", "/var/log/go/gotool.log")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}

	flag.Parse()

	if *file == "" || *tweakTime == 0.0 {
		flag.Usage()

		os.Exit(1)
		return
	}
}

func main() {
	fmt.Println(file, tweakTime)
	err := readSrtFile()
	if err != nil {
		panic(err)
	}

	err = copyFile()
	if err != nil {
		panic(err)
	}
}
