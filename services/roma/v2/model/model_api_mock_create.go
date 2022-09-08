package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// mock后端详情
type ApiMockCreate struct {

	// 描述信息。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 返回结果
	ResultContent *string `json:"result_content,omitempty"`

	// 版本。字符长度不超过64
	Version *string `json:"version,omitempty"`

	// 后端自定义认证ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// mock后端自定义状态码： \"200\": \"OK\", \"201\": \"Created\", \"202\": \"Accepted\", \"203\": \"NonAuthoritativeInformation\", \"204\": \"NoContent\", \"205\": \"ResetContent\", \"206\": \"PartialContent\", \"300\": \"MultipleChoices\", \"301\": \"MovedPermanently\", \"302\": \"Found\", \"303\": \"SeeOther\", \"304\": \"NotModified\", \"305\": \"UseProxy\", \"306\": \"Unused\", \"307\": \"TemporaryRedirect\", \"400\": \"BadRequest\", \"401\": \"Unauthorized\", \"402\": \"PaymentRequired\", \"403\": \"Forbidden\", \"404\": \"NotFound\", \"405\": \"MethodNotAllowed\", \"406\": \"NotAcceptable\", \"407\": \"ProxyAuthenticationRequired\", \"408\": \"RequestTimeout\", \"409\": \"Conflict\", \"410\": \"Gone\", \"411\": \"LengthRequired\", \"412\": \"PreconditionFailed\", \"413\": \"RequestEntityTooLarge\", \"414\": \"RequestURITooLong\", \"415\": \"UnsupportedMediaType\", \"416\": \"RequestedRangeNotSatisfiable\", \"417\": \"ExpectationFailed\", \"450\": \"ParameterRequried\", \"451\": \"MethodConnectException\", \"500\": \"InternalServerError\", \"501\": \"NotImplemented\", \"502\": \"BadGateway\", \"503\": \"ServiceUnavailable\", \"504\": \"GatewayTimeout\", \"505\": \"HTTPVersionNotSupported\",
	StatusCode *ApiMockCreateStatusCode `json:"status_code,omitempty"`

	// mock后端自定义响应头header
	Header *[]MockApiBaseInfoHeader `json:"header,omitempty"`
}

func (o ApiMockCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiMockCreate struct{}"
	}

	return strings.Join([]string{"ApiMockCreate", string(data)}, " ")
}

type ApiMockCreateStatusCode struct {
	value int32
}

type ApiMockCreateStatusCodeEnum struct {
	E_200 ApiMockCreateStatusCode
	E_201 ApiMockCreateStatusCode
	E_202 ApiMockCreateStatusCode
	E_203 ApiMockCreateStatusCode
	E_204 ApiMockCreateStatusCode
	E_205 ApiMockCreateStatusCode
	E_206 ApiMockCreateStatusCode
	E_300 ApiMockCreateStatusCode
	E_301 ApiMockCreateStatusCode
	E_302 ApiMockCreateStatusCode
	E_303 ApiMockCreateStatusCode
	E_304 ApiMockCreateStatusCode
	E_305 ApiMockCreateStatusCode
	E_306 ApiMockCreateStatusCode
	E_307 ApiMockCreateStatusCode
	E_400 ApiMockCreateStatusCode
	E_401 ApiMockCreateStatusCode
	E_402 ApiMockCreateStatusCode
	E_403 ApiMockCreateStatusCode
	E_404 ApiMockCreateStatusCode
	E_405 ApiMockCreateStatusCode
	E_406 ApiMockCreateStatusCode
	E_407 ApiMockCreateStatusCode
	E_408 ApiMockCreateStatusCode
	E_409 ApiMockCreateStatusCode
	E_410 ApiMockCreateStatusCode
	E_411 ApiMockCreateStatusCode
	E_412 ApiMockCreateStatusCode
	E_413 ApiMockCreateStatusCode
	E_414 ApiMockCreateStatusCode
	E_415 ApiMockCreateStatusCode
	E_416 ApiMockCreateStatusCode
	E_417 ApiMockCreateStatusCode
	E_450 ApiMockCreateStatusCode
	E_451 ApiMockCreateStatusCode
	E_500 ApiMockCreateStatusCode
	E_501 ApiMockCreateStatusCode
	E_502 ApiMockCreateStatusCode
	E_503 ApiMockCreateStatusCode
	E_504 ApiMockCreateStatusCode
	E_505 ApiMockCreateStatusCode
}

func GetApiMockCreateStatusCodeEnum() ApiMockCreateStatusCodeEnum {
	return ApiMockCreateStatusCodeEnum{
		E_200: ApiMockCreateStatusCode{
			value: 200,
		}, E_201: ApiMockCreateStatusCode{
			value: 201,
		}, E_202: ApiMockCreateStatusCode{
			value: 202,
		}, E_203: ApiMockCreateStatusCode{
			value: 203,
		}, E_204: ApiMockCreateStatusCode{
			value: 204,
		}, E_205: ApiMockCreateStatusCode{
			value: 205,
		}, E_206: ApiMockCreateStatusCode{
			value: 206,
		}, E_300: ApiMockCreateStatusCode{
			value: 300,
		}, E_301: ApiMockCreateStatusCode{
			value: 301,
		}, E_302: ApiMockCreateStatusCode{
			value: 302,
		}, E_303: ApiMockCreateStatusCode{
			value: 303,
		}, E_304: ApiMockCreateStatusCode{
			value: 304,
		}, E_305: ApiMockCreateStatusCode{
			value: 305,
		}, E_306: ApiMockCreateStatusCode{
			value: 306,
		}, E_307: ApiMockCreateStatusCode{
			value: 307,
		}, E_400: ApiMockCreateStatusCode{
			value: 400,
		}, E_401: ApiMockCreateStatusCode{
			value: 401,
		}, E_402: ApiMockCreateStatusCode{
			value: 402,
		}, E_403: ApiMockCreateStatusCode{
			value: 403,
		}, E_404: ApiMockCreateStatusCode{
			value: 404,
		}, E_405: ApiMockCreateStatusCode{
			value: 405,
		}, E_406: ApiMockCreateStatusCode{
			value: 406,
		}, E_407: ApiMockCreateStatusCode{
			value: 407,
		}, E_408: ApiMockCreateStatusCode{
			value: 408,
		}, E_409: ApiMockCreateStatusCode{
			value: 409,
		}, E_410: ApiMockCreateStatusCode{
			value: 410,
		}, E_411: ApiMockCreateStatusCode{
			value: 411,
		}, E_412: ApiMockCreateStatusCode{
			value: 412,
		}, E_413: ApiMockCreateStatusCode{
			value: 413,
		}, E_414: ApiMockCreateStatusCode{
			value: 414,
		}, E_415: ApiMockCreateStatusCode{
			value: 415,
		}, E_416: ApiMockCreateStatusCode{
			value: 416,
		}, E_417: ApiMockCreateStatusCode{
			value: 417,
		}, E_450: ApiMockCreateStatusCode{
			value: 450,
		}, E_451: ApiMockCreateStatusCode{
			value: 451,
		}, E_500: ApiMockCreateStatusCode{
			value: 500,
		}, E_501: ApiMockCreateStatusCode{
			value: 501,
		}, E_502: ApiMockCreateStatusCode{
			value: 502,
		}, E_503: ApiMockCreateStatusCode{
			value: 503,
		}, E_504: ApiMockCreateStatusCode{
			value: 504,
		}, E_505: ApiMockCreateStatusCode{
			value: 505,
		},
	}
}

func (c ApiMockCreateStatusCode) Value() int32 {
	return c.value
}

func (c ApiMockCreateStatusCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiMockCreateStatusCode) UnmarshalJSON(b []byte) error {
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
