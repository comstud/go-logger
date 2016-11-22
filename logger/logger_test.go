package logger

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

// Returns message logged, if any, in callback fn
func doLog(prefix string, fn func(Logger)) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	fn(NewDefaultLogger(w, prefix))
	w.Flush()
	byt, _ := ioutil.ReadAll(bufio.NewReader(&b))
	return string(byt)
}

func TestHeader(t *testing.T) {
	line := doLog("", func(l Logger) {
		l.LogWarn()
	})
	parts := strings.SplitN(line, " ", 3)
	if len(parts) != 3 {
		t.Errorf("Didn't find 3 parts of log msg: %s", line)
	}

	m, _ := regexp.MatchString(`^\d{4}/\d{2}/\d{2}$`, parts[0])
	if !m {
		t.Errorf("Date not correct at start: %s", parts[0])
	}

	m, _ = regexp.MatchString(`^\d{2}:\d{2}:\d{2}\.\d+$`, parts[1])
	if !m {
		t.Errorf("Time not correct: %s", parts[1])
	}
}

func TestPrefix(t *testing.T) {
	line := doLog("myprefix ", func(l Logger) {
		l.LogWarn()
	})

	if !strings.HasPrefix(line, "myprefix ") {
		t.Error("Log message doesn't start with correct prefix")
	}

	parts := strings.SplitN(line, " ", 4)
	if len(parts) != 4 {
		t.Errorf("Didn't find 4 parts of log msg: %s", line)
	}

	m, _ := regexp.MatchString(`^\d{4}/\d{2}/\d{2}$`, parts[1])
	if !m {
		t.Errorf("Date not correct at start: %s", parts[1])
	}

	m, _ = regexp.MatchString(`^\d{2}:\d{2}:\d{2}\.\d+$`, parts[2])
	if !m {
		t.Errorf("Time not correct: %s", parts[2])
	}
}

func TestDebug(t *testing.T) {
	line := doLog("", func(l Logger) {
		l.LogDebug("1", 2, 3)
	})

	parts := strings.SplitN(line, " ", 4)

	if parts[2] != "[DEBUG]" {
		t.Errorf("No [DEBUG] found: %s", parts[2])
	}

	if parts[3] != "1 2 3\n" {
		t.Errorf("Log part incorrect: %s", parts[3])
	}
}

func TestError(t *testing.T) {
	line := doLog("", func(l Logger) {
		l.LogError("1", 2, 3)
	})

	parts := strings.SplitN(line, " ", 4)

	if parts[2] != "[ERROR]" {
		t.Errorf("No [ERROR] found: %s", parts[2])
	}

	if parts[3] != "1 2 3\n" {
		t.Errorf("Log part incorrect: %s", parts[3])
	}
}

func TestInfo(t *testing.T) {
	line := doLog("", func(l Logger) {
		l.LogInfo("1", 2, 3)
	})

	parts := strings.SplitN(line, " ", 4)

	if parts[2] != "[INFO]" {
		t.Errorf("No [INFO] found: %s", parts[2])
	}

	if parts[3] != "1 2 3\n" {
		t.Errorf("Log part incorrect: %s", parts[3])
	}
}

func TestWarn(t *testing.T) {
	line := doLog("", func(l Logger) {
		l.LogWarn("1", 2, 3)
	})

	parts := strings.SplitN(line, " ", 4)

	if parts[2] != "[WARN]" {
		t.Errorf("No [WARN] found: %s", parts[2])
	}

	if parts[3] != "1 2 3\n" {
		t.Errorf("Log part incorrect: %s", parts[3])
	}
}
