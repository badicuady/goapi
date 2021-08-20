package common

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"runtime"
)

func Trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return file, line, "?"
	}

	return file, line, fn.Name()
}

func ReadJson(data interface{}, reader io.ReadCloser) error {
	body, err1 := ioutil.ReadAll(reader)
	if err1 != nil {
		return err1
	}

	err2 := json.Unmarshal(body, &data)
	if err2 != nil {
		return err2
	}

	return nil
}
