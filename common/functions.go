package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func ReadJson(data interface{}, w http.ResponseWriter, r *http.Request) error {
	body, err1 := ioutil.ReadAll(r.Body)
	if err1 != nil {
		return err1
	}

	err2 := json.Unmarshal(body, &data)
	if err2 != nil {
		return err2
	}

	return nil
}

func WriteJson(v interface{}, w http.ResponseWriter) error {
	return WriteJsonWithStatusCode(v, 200, w)
}

func WriteJsonWithStatusCode(v interface{}, statusCode int, w http.ResponseWriter) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(b)

	return nil
}

func WriteError(err error, w http.ResponseWriter) {
	WriteErrorWithStatusCode(err, http.StatusInternalServerError, w)
}

func WriteErrorWithStatusCode(err error, statusCode int, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
