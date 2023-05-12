package core

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Body(r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	for k := range r.Form {
		data[k] = r.FormValue(k)
	}
	if len(data) == 0 {
		_ = json.NewDecoder(r.Body).Decode(&data)
	}
	return data
}

func Header(r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	for k := range r.Header {
		data[k] = r.Header.Get(k)
	}
	return data
}

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
		v := r.FormValue(k)
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
		v := r.FormValue(k)
		if v != "" {
			data[k] = v
		}
	}
	return data
}

func ParseOrigins(originAllowed string) []string {
	origins := []string{}
	if originAllowed != "" {
		originList := strings.Split(originAllowed, ",")
		for _, origin := range originList {
			trimmedOrigin := strings.TrimSpace(origin)
			origins = append(origins, trimmedOrigin)
		}
	}
	return origins
}
