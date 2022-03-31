package barrier

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers",
			"http://httpbin.org/user-agent"}

		result := captureBarrierOutPut(endpoints...)
		if !strings.Contains(result, "Accept-Encoding") ||
			!strings.Contains(result, "user-agent") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url",
			"http://httpbin.org/user-agent"}
		result := captureBarrierOutPut(endpoints...)
		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers",
			"http://httpbin.org/user-agent"}
		timeoutMilliseconds = 1
		result := captureBarrierOutPut(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})

}

func captureBarrierOutPut(enpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer
	out := make(chan string)
	go func() {
		// var buf bytes.Buffer
		// io.Copy(&buf, reader)
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		out <- buf.String()
	}()

	barrier(enpoints...)

	writer.Close()
	temp := <-out

	return temp
}
