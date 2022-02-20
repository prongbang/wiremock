package core

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func Body(c *fiber.Ctx) map[string]interface{} {
	body := map[string]interface{}{}
	b := c.Body()
	_ = json.Unmarshal(b, &body)
	return body
}

func BodyDecode(c *fiber.Ctx, v interface{}) {
	reader := bytes.NewReader(c.Body())
	_ = json.NewDecoder(reader).Decode(v)
}

func BindHeader(mockHeader map[string]interface{}, c *fiber.Ctx) map[string]interface{} {
	data := map[string]interface{}{}
	for k := range mockHeader {
		header := c.GetReqHeaders()
		v := header[k]
		if v != "" {
			data[k] = v
		}
	}
	return data
}

func BindBody(mockBody map[string]interface{}, c *fiber.Ctx) map[string]interface{} {
	data := map[string]interface{}{}
	body := Body(c)
	for k := range mockBody {
		v := body[k]
		if v != "" {
			data[k] = v
		}
	}
	if len(data) == 0 {
		BodyDecode(c, &data)
	}
	return data
}

func BindCaseBody(mockBody map[string]interface{}, c *fiber.Ctx) map[string]interface{} {
	data := map[string]interface{}{}
	body := fiber.Map{}
	_ = c.BodyParser(&body)
	for k := range mockBody {
		v := body[k]
		if v != "" {
			data[k] = v
		}
	}
	return data
}
