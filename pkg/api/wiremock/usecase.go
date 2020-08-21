package wiremock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prongbang/wiremock/pkg/api/core"
	"github.com/prongbang/wiremock/pkg/config"
	"github.com/prongbang/wiremock/pkg/status"
)

type UseCase interface {
	CasesMatching(r *http.Request, path string, cases map[string]Cases, params Parameters) CaseMatching
	ParameterMatching(params Parameters) Matching
	GetMockResponse(resp Response) []byte
	ReadSourceRouteYml(routeName string) []byte
}

type useCase struct {
}

func (u *useCase) CasesMatching(r *http.Request, path string, cases map[string]Cases, params Parameters) CaseMatching {

	// Get request
	body := map[string]interface{}{}
	_ = json.NewDecoder(r.Body).Decode(&body)

	// Process header matching
	require := map[string]interface{}{}
	errors := map[string]interface{}{}
	matchingHeader := 0
	for k, v := range params.ReqBody.Mock {
		vs := fmt.Sprintf("%v", v)
		ks := fmt.Sprintf("%v", params.ReqHeader.Http[k])
		if vs == ks {
			matchingHeader = matchingHeader + 1
			continue
		}
		if params.ReqHeader.Http[k] == nil {
			errors[k] = "Require header " + k
		} else {
			errors[k] = "The header " + k + " not match"
		}
	}
	if len(errors) > 0 {
		require["errors"] = errors
	}
	require["message"] = "validation error"
	require["status"] = "error"
	result, err := json.Marshal(require)
	if err != nil {
		result = []byte("{}")
	}
	matchingHeaderRequest := len(params.ReqBody.Mock) == matchingHeader

	// Process body matching
	matchingBodyRequest := false
	var foundCase Cases

	for _, vMock := range cases {
		matchingBody := 0
		vMock.Response.FileName = path
		if len(body) == 0 {
			body = core.BindCaseBody(vMock.Body, r)
		}
		for ck, cv := range vMock.Body {
			vs := fmt.Sprintf("%v", cv)
			ks := fmt.Sprintf("%v", body[ck])

			// Check require field value is not empty
			if vs == "*" {
				if body[ck] != nil {
					matchingBody = matchingBody + 1
				}
			}

			// Value matching
			if vs == ks {
				matchingBody = matchingBody + 1
			}
		}

		// Contains value
		matchingBodyRequest = len(vMock.Body) == matchingBody
		if matchingBodyRequest {
			foundCase = vMock
			break
		}
	}

	return CaseMatching{
		IsMatch: matchingBodyRequest && matchingHeaderRequest,
		Result:  result,
		Case:    foundCase,
	}
}

func (u *useCase) ParameterMatching(params Parameters) Matching {
	require := map[string]interface{}{}
	errors := map[string]interface{}{}
	matchingHeader := 0
	matchingBody := 0
	for k, v := range params.ReqBody.Mock {
		vs := fmt.Sprintf("%v", v)
		ks := fmt.Sprintf("%v", params.ReqBody.Http[k])
		if vs == ks {
			matchingBody = matchingBody + 1
			continue
		}
		if params.ReqBody.Http[k] == nil {
			errors[k] = "Require " + k
		} else {
			errors[k] = "The " + k + " not match"
		}
	}

	for k, v := range params.ReqHeader.Mock {
		vs := fmt.Sprintf("%v", v)
		ks := fmt.Sprintf("%v", params.ReqHeader.Http[k])
		if vs == ks {
			matchingHeader = matchingHeader + 1
			continue
		}
		if params.ReqHeader.Http[k] == nil {
			errors[k] = "Require header " + k
		} else {
			errors[k] = "The header " + k + " not match"
		}
	}

	if len(errors) > 0 {
		require["errors"] = errors
		require["message"] = "validation error"
		require["status"] = "error"
	}

	result, err := json.Marshal(require)
	if err != nil {
		result = []byte("{}")
	}

	isMatchHeader := len(params.ReqHeader.Mock) == matchingHeader
	isMatchBody := len(params.ReqBody.Mock) == matchingBody

	return Matching{
		Result:  result,
		IsMatch: isMatchBody && isMatchHeader,
	}
}

func (u *useCase) GetMockResponse(resp Response) []byte {
	if resp.BodyFile != "" {
		bodyFile := fmt.Sprintf(config.MockResponsePath, resp.FileName, resp.BodyFile)
		source, err := ioutil.ReadFile(bodyFile)
		if err != nil {
			return []byte("{}")
		}
		return source
	}
	return []byte(resp.Body)
}

func (u *useCase) ReadSourceRouteYml(routeName string) []byte {
	pattern := status.Pattern()
	filename := fmt.Sprintf(config.MockRouteYmlPath, routeName)
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(pattern)
	}
	return source
}

func NewUseCase() UseCase {
	return &useCase{}
}
