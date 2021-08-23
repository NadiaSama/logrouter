package logrouter

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func TestLevelMapper(t *testing.T) {
	mapper := NewLevelMapper()

	buf := bytes.NewBuffer(make([]byte, 0))
	logger := log.NewLogfmtLogger(buf)
	mapper.AddInfo(logger)

	level.Info(mapper).Log("message", "a")

	raw, _ := ioutil.ReadAll(buf)
	if string(raw) != "level=info message=a\n" {
		t.Errorf("bad buf got='%s'", string(raw))
	}
}
