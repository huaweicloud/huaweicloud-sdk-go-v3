package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyMockBase struct {

	// 返回结果
	ResultContent *string `json:"result_content,omitempty" xml:"result_content"`

	// mock后端自定义状态码： \"200\": \"OK\", \"201\": \"Created\", \"202\": \"Accepted\", \"203\": \"NonAuthoritativeInformation\", \"204\": \"NoContent\", \"205\": \"ResetContent\", \"206\": \"PartialContent\", \"300\": \"MultipleChoices\", \"301\": \"MovedPermanently\", \"302\": \"Found\", \"303\": \"SeeOther\", \"304\": \"NotModified\", \"305\": \"UseProxy\", \"306\": \"Unused\", \"307\": \"TemporaryRedirect\", \"400\": \"BadRequest\", \"401\": \"Unauthorized\", \"402\": \"PaymentRequired\", \"403\": \"Forbidden\", \"404\": \"NotFound\", \"405\": \"MethodNotAllowed\", \"406\": \"NotAcceptable\", \"407\": \"ProxyAuthenticationRequired\", \"408\": \"RequestTimeout\", \"409\": \"Conflict\", \"410\": \"Gone\", \"411\": \"LengthRequired\", \"412\": \"PreconditionFailed\", \"413\": \"RequestEntityTooLarge\", \"414\": \"RequestURITooLong\", \"415\": \"UnsupportedMediaType\", \"416\": \"RequestedRangeNotSatisfiable\", \"417\": \"ExpectationFailed\", \"450\": \"ParameterRequried\", \"451\": \"MethodConnectException\", \"500\": \"InternalServerError\", \"501\": \"NotImplemented\", \"502\": \"BadGateway\", \"503\": \"ServiceUnavailable\", \"504\": \"GatewayTimeout\", \"505\": \"HTTPVersionNotSupported\",
	StatusCode *ApiPolicyMockBaseStatusCode `json:"status_code,omitempty" xml:"status_code"`

	// mock后端自定义响应头header
	Header *[]MockApiBaseInfoHeader `json:"header,omitempty" xml:"header"`
}

func (o ApiPolicyMockBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyMockBase struct{}"
	}

	return strings.Join([]string{"ApiPolicyMockBase", string(data)}, " ")
}

type ApiPolicyMockBaseStatusCode struct {
	value int32
}

type ApiPolicyMockBaseStatusCodeEnum struct {
	E_200 ApiPolicyMockBaseStatusCode
	E_201 ApiPolicyMockBaseStatusCode
	E_202 ApiPolicyMockBaseStatusCode
	E_203 ApiPolicyMockBaseStatusCode
	E_204 ApiPolicyMockBaseStatusCode
	E_205 ApiPolicyMockBaseStatusCode
	E_206 ApiPolicyMockBaseStatusCode
	E_300 ApiPolicyMockBaseStatusCode
	E_301 ApiPolicyMockBaseStatusCode
	E_302 ApiPolicyMockBaseStatusCode
	E_303 ApiPolicyMockBaseStatusCode
	E_304 ApiPolicyMockBaseStatusCode
	E_305 ApiPolicyMockBaseStatusCode
	E_306 ApiPolicyMockBaseStatusCode
	E_307 ApiPolicyMockBaseStatusCode
	E_400 ApiPolicyMockBaseStatusCode
	E_401 ApiPolicyMockBaseStatusCode
	E_402 ApiPolicyMockBaseStatusCode
	E_403 ApiPolicyMockBaseStatusCode
	E_404 ApiPolicyMockBaseStatusCode
	E_405 ApiPolicyMockBaseStatusCode
	E_406 ApiPolicyMockBaseStatusCode
	E_407 ApiPolicyMockBaseStatusCode
	E_408 ApiPolicyMockBaseStatusCode
	E_409 ApiPolicyMockBaseStatusCode
	E_410 ApiPolicyMockBaseStatusCode
	E_411 ApiPolicyMockBaseStatusCode
	E_412 ApiPolicyMockBaseStatusCode
	E_413 ApiPolicyMockBaseStatusCode
	E_414 ApiPolicyMockBaseStatusCode
	E_415 ApiPolicyMockBaseStatusCode
	E_416 ApiPolicyMockBaseStatusCode
	E_417 ApiPolicyMockBaseStatusCode
	E_450 ApiPolicyMockBaseStatusCode
	E_451 ApiPolicyMockBaseStatusCode
	E_500 ApiPolicyMockBaseStatusCode
	E_501 ApiPolicyMockBaseStatusCode
	E_502 ApiPolicyMockBaseStatusCode
	E_503 ApiPolicyMockBaseStatusCode
	E_504 ApiPolicyMockBaseStatusCode
	E_505 ApiPolicyMockBaseStatusCode
}

func GetApiPolicyMockBaseStatusCodeEnum() ApiPolicyMockBaseStatusCodeEnum {
	return ApiPolicyMockBaseStatusCodeEnum{
		E_200: ApiPolicyMockBaseStatusCode{
			value: 200,
		}, E_201: ApiPolicyMockBaseStatusCode{
			value: 201,
		}, E_202: ApiPolicyMockBaseStatusCode{
			value: 202,
		}, E_203: ApiPolicyMockBaseStatusCode{
			value: 203,
		}, E_204: ApiPolicyMockBaseStatusCode{
			value: 204,
		}, E_205: ApiPolicyMockBaseStatusCode{
			value: 205,
		}, E_206: ApiPolicyMockBaseStatusCode{
			value: 206,
		}, E_300: ApiPolicyMockBaseStatusCode{
			value: 300,
		}, E_301: ApiPolicyMockBaseStatusCode{
			value: 301,
		}, E_302: ApiPolicyMockBaseStatusCode{
			value: 302,
		}, E_303: ApiPolicyMockBaseStatusCode{
			value: 303,
		}, E_304: ApiPolicyMockBaseStatusCode{
			value: 304,
		}, E_305: ApiPolicyMockBaseStatusCode{
			value: 305,
		}, E_306: ApiPolicyMockBaseStatusCode{
			value: 306,
		}, E_307: ApiPolicyMockBaseStatusCode{
			value: 307,
		}, E_400: ApiPolicyMockBaseStatusCode{
			value: 400,
		}, E_401: ApiPolicyMockBaseStatusCode{
			value: 401,
		}, E_402: ApiPolicyMockBaseStatusCode{
			value: 402,
		}, E_403: ApiPolicyMockBaseStatusCode{
			value: 403,
		}, E_404: ApiPolicyMockBaseStatusCode{
			value: 404,
		}, E_405: ApiPolicyMockBaseStatusCode{
			value: 405,
		}, E_406: ApiPolicyMockBaseStatusCode{
			value: 406,
		}, E_407: ApiPolicyMockBaseStatusCode{
			value: 407,
		}, E_408: ApiPolicyMockBaseStatusCode{
			value: 408,
		}, E_409: ApiPolicyMockBaseStatusCode{
			value: 409,
		}, E_410: ApiPolicyMockBaseStatusCode{
			value: 410,
		}, E_411: ApiPolicyMockBaseStatusCode{
			value: 411,
		}, E_412: ApiPolicyMockBaseStatusCode{
			value: 412,
		}, E_413: ApiPolicyMockBaseStatusCode{
			value: 413,
		}, E_414: ApiPolicyMockBaseStatusCode{
			value: 414,
		}, E_415: ApiPolicyMockBaseStatusCode{
			value: 415,
		}, E_416: ApiPolicyMockBaseStatusCode{
			value: 416,
		}, E_417: ApiPolicyMockBaseStatusCode{
			value: 417,
		}, E_450: ApiPolicyMockBaseStatusCode{
			value: 450,
		}, E_451: ApiPolicyMockBaseStatusCode{
			value: 451,
		}, E_500: ApiPolicyMockBaseStatusCode{
			value: 500,
		}, E_501: ApiPolicyMockBaseStatusCode{
			value: 501,
		}, E_502: ApiPolicyMockBaseStatusCode{
			value: 502,
		}, E_503: ApiPolicyMockBaseStatusCode{
			value: 503,
		}, E_504: ApiPolicyMockBaseStatusCode{
			value: 504,
		}, E_505: ApiPolicyMockBaseStatusCode{
			value: 505,
		},
	}
}

func (c ApiPolicyMockBaseStatusCode) Value() int32 {
	return c.value
}

func (c ApiPolicyMockBaseStatusCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyMockBaseStatusCode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
