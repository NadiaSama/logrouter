package logrouter

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/go-kit/log"
)

func TestLogger(t *testing.T) {
	var (
		key  interface{} = "key"
		val1 interface{} = "val1"
		val2 interface{} = "val2"
	)
	expectBuf := func(b io.Reader, expect string, t *testing.T) {
		have, _ := ioutil.ReadAll(b)
		if string(have) != expect {
			t.Errorf("test buf fail expect=%s have='%s'", expect, have)
		}
	}

	t.Run("test log router", func(t *testing.T) {
		buf1 := bytes.NewBuffer(make([]byte, 0))
		buf2 := bytes.NewBuffer(make([]byte, 0))
		buf3 := bytes.NewBuffer(make([]byte, 0))
		buf4 := bytes.NewBuffer(make([]byte, 0))

		logger1 := log.NewLogfmtLogger(buf1)
		logger2 := log.NewLogfmtLogger(buf2)
		logger3 := log.NewLogfmtLogger(buf3)
		logger4 := log.NewLogfmtLogger(buf4)

		mapper := NewMapper(key)
		mapper.AddLogger(val1, logger1, logger2)
		mapper.AddLogger(val2, logger3)

		if err := mapper.Log(key, "val3"); err != ErrNoLogger {
			t.Errorf("test no logger failed err %v", err)
		}

		mapper.Log(key, val1)
		mapper.Log(key, val2)
		mapper.SetDefault(logger4)

		mapper.Log(key, "val3")

		expectBuf(buf1, "key=val1\n", t)
		expectBuf(buf2, "key=val1\n", t)
		expectBuf(buf3, "key=val2\n", t)
		expectBuf(buf4, "key=val3\n", t)

	})

	t.Run("test log odd element", func(t *testing.T) {
		buf := bytes.NewBuffer(make([]byte, 0))
		logger := log.NewLogfmtLogger(buf)
		mapper := NewMapper(key)
		mapper.AddLogger(val1, logger)

		if err := mapper.Log("a", "1", "b", "2", key); err != ErrNoLogger {
			t.Errorf("expect err got %v", err)
		}

		mapper.SetDefault(logger)
		mapper.Log("a", 1, "b", 2, key, "3")
		expectBuf(buf, "a=1 b=2 key=3\n", t)
	})
}
