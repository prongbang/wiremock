package wiremock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/prongbang/wiremock/pkg/config"
	"github.com/prongbang/wiremock/pkg/status"
)

type UseCase interface {
	ParameterMatching(params Parameters) Matching
	GetMockResponse(resp Response) []byte
	ReadSourceRouteYml(routeName string) []byte
}

type useCase struct {
}

func (u *useCase) ParameterMatching(params Parameters) Matching {
	require := map[string]interface{}{}
	errors := map[string]interface{}{}
	matchingHeader := 0
	matchingBody := 0
	for k, v := range params.MockReqBody {
		vs := fmt.Sprintf("%v", v)
		ks := fmt.Sprintf("%v", params.HttpReqBody[k])
		if vs == ks {
			matchingBody = matchingBody + 1
			continue
		}
		if params.HttpReqBody[k] == nil {
			errors[k] = "Require " + k
		} else {
			errors[k] = "The " + k + " not match"
		}
	}

	for k, v := range params.MockReqHeader {
		vs := fmt.Sprintf("%v", v)
		ks := fmt.Sprintf("%v", params.HttpReqHeader[k])
		if vs == ks {
			matchingHeader = matchingHeader + 1
			continue
		}
		if params.HttpReqHeader[k] == nil {
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

	isMatchHeader := len(params.MockReqHeader) == matchingHeader
	isMatchBody := len(params.MockReqBody) == matchingBody

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
