package v3

import (
	"encoding/json"
	"strconv"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/response"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
)

type ErrorHandler interface {
	HandleError(req *request.DefaultHttpRequest, resp *response.DefaultHttpResponse) error
}

type CceErrorHandler struct{}

type errMap map[string]interface{}

func (m errMap) getStringValue(key string) string {
	var result string

	value, isExist := m[key]
	if !isExist {
		return result
	}

	if strVal, ok := value.(string); ok {
		result = strVal
	}

	return result
}

func (h CceErrorHandler) HandleError(req *request.DefaultHttpRequest, resp *response.DefaultHttpResponse) error {
	if resp.GetStatusCode() < 400 {
		return nil
	}

	sr := &sdkerr.ServiceResponseError{
		StatusCode: resp.GetStatusCode(),
		RequestId:  resp.GetHeader("X-Request-Id"),
	}

	var dataMap errMap
	data := []byte(resp.GetBody())
	if err := json.Unmarshal(data, &dataMap); err != nil {
		sr.ErrorCode = strconv.Itoa(resp.GetStatusCode())
		sr.ErrorMessage = string(data)
		return err
	}
	sr.ErrorCode = dataMap.getStringValue("error_code")
	if sr.ErrorCode == "" {
		sr.ErrorCode = dataMap.getStringValue("errorCode")
	}
	sr.ErrorMessage = dataMap.getStringValue("error_msg")
	if sr.ErrorMessage == "" {
		sr.ErrorMessage = dataMap.getStringValue("errorMessage")
	}
	sr.EncodedAuthorizationMessage = dataMap.getStringValue("encoded_authorization_message")
	message := dataMap.getStringValue("message")
	if message != "" && message != sr.ErrorMessage {
		if sr.ErrorMessage == "" {
			sr.ErrorMessage = message
		} else {
			sr.ErrorMessage += ", " + message
		}
	}
	return sr
}
