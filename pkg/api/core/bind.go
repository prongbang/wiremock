package core

import (
	"encoding/json"
	"net/http"
)

func BindHeader(mockHeader map[string]interface{}, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	for k := range mockHeader {
		v := r.Header.Get(k)
		if v != "" {
			data[k] = v
		}
	}
	return data
}

func BindBody(mockBody map[string]interface{}, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	for k := range mockBody {
		v := r.PostFormValue(k)
		if v != "" {
			data[k] = v
		}
	}
	if len(data) == 0 {
		_ = json.NewDecoder(r.Body).Decode(&data)
	}
	return data
}

func BindCaseBody(mockBody map[string]interface{}, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	for k := range mockBody {
		v := r.PostFormValue(k)
		if v != "" {
			data[k] = v
		}
	}
	return data
}
