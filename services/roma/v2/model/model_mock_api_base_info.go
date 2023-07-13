package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MockApiBaseInfo mock后端详情
type MockApiBaseInfo struct {

	// 描述信息。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 返回结果
	ResultContent *string `json:"result_content,omitempty"`

	// 版本。字符长度不超过64
	Version *string `json:"version,omitempty"`

	// 后端自定义认证ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// mock后端自定义状态码： \"200\": \"OK\", \"201\": \"Created\", \"202\": \"Accepted\", \"203\": \"NonAuthoritativeInformation\", \"204\": \"NoContent\", \"205\": \"ResetContent\", \"206\": \"PartialContent\", \"300\": \"MultipleChoices\", \"301\": \"MovedPermanently\", \"302\": \"Found\", \"303\": \"SeeOther\", \"304\": \"NotModified\", \"305\": \"UseProxy\", \"306\": \"Unused\", \"307\": \"TemporaryRedirect\", \"400\": \"BadRequest\", \"401\": \"Unauthorized\", \"402\": \"PaymentRequired\", \"403\": \"Forbidden\", \"404\": \"NotFound\", \"405\": \"MethodNotAllowed\", \"406\": \"NotAcceptable\", \"407\": \"ProxyAuthenticationRequired\", \"408\": \"RequestTimeout\", \"409\": \"Conflict\", \"410\": \"Gone\", \"411\": \"LengthRequired\", \"412\": \"PreconditionFailed\", \"413\": \"RequestEntityTooLarge\", \"414\": \"RequestURITooLong\", \"415\": \"UnsupportedMediaType\", \"416\": \"RequestedRangeNotSatisfiable\", \"417\": \"ExpectationFailed\", \"450\": \"ParameterRequried\", \"451\": \"MethodConnectException\", \"500\": \"InternalServerError\", \"501\": \"NotImplemented\", \"502\": \"BadGateway\", \"503\": \"ServiceUnavailable\", \"504\": \"GatewayTimeout\", \"505\": \"HTTPVersionNotSupported\",
	StatusCode *MockApiBaseInfoStatusCode `json:"status_code,omitempty"`

	// mock后端自定义响应头header  格式：[{\"key\":\"\",\"value\": \"\", \"remark:\"\"}, {\"key2\":\"\",\"value2\": \"\",\"remark2:\"\"}]  参数说明：  key：mock后端自定义响应头header key, 支持英文，数字，中划线，且只能以英文字母或数字开头，1 ~ 64字符  value： mock后端自定义响应头header value，中文字符必须为UTF-8或者unicode编码, 不能为空，最大长度为10240  remark：mock后端自定义响应头header remark，中文字符必须为UTF-8 或者unicode编码，可以为空，最大长度为2048
	Header *string `json:"header,omitempty"`
}

func (o MockApiBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MockApiBaseInfo struct{}"
	}

	return strings.Join([]string{"MockApiBaseInfo", string(data)}, " ")
}

type MockApiBaseInfoStatusCode struct {
	value int32
}

type MockApiBaseInfoStatusCodeEnum struct {
	E_200 MockApiBaseInfoStatusCode
	E_201 MockApiBaseInfoStatusCode
	E_202 MockApiBaseInfoStatusCode
	E_203 MockApiBaseInfoStatusCode
	E_204 MockApiBaseInfoStatusCode
	E_205 MockApiBaseInfoStatusCode
	E_206 MockApiBaseInfoStatusCode
	E_300 MockApiBaseInfoStatusCode
	E_301 MockApiBaseInfoStatusCode
	E_302 MockApiBaseInfoStatusCode
	E_303 MockApiBaseInfoStatusCode
	E_304 MockApiBaseInfoStatusCode
	E_305 MockApiBaseInfoStatusCode
	E_306 MockApiBaseInfoStatusCode
	E_307 MockApiBaseInfoStatusCode
	E_400 MockApiBaseInfoStatusCode
	E_401 MockApiBaseInfoStatusCode
	E_402 MockApiBaseInfoStatusCode
	E_403 MockApiBaseInfoStatusCode
	E_404 MockApiBaseInfoStatusCode
	E_405 MockApiBaseInfoStatusCode
	E_406 MockApiBaseInfoStatusCode
	E_407 MockApiBaseInfoStatusCode
	E_408 MockApiBaseInfoStatusCode
	E_409 MockApiBaseInfoStatusCode
	E_410 MockApiBaseInfoStatusCode
	E_411 MockApiBaseInfoStatusCode
	E_412 MockApiBaseInfoStatusCode
	E_413 MockApiBaseInfoStatusCode
	E_414 MockApiBaseInfoStatusCode
	E_415 MockApiBaseInfoStatusCode
	E_416 MockApiBaseInfoStatusCode
	E_417 MockApiBaseInfoStatusCode
	E_450 MockApiBaseInfoStatusCode
	E_451 MockApiBaseInfoStatusCode
	E_500 MockApiBaseInfoStatusCode
	E_501 MockApiBaseInfoStatusCode
	E_502 MockApiBaseInfoStatusCode
	E_503 MockApiBaseInfoStatusCode
	E_504 MockApiBaseInfoStatusCode
	E_505 MockApiBaseInfoStatusCode
}

func GetMockApiBaseInfoStatusCodeEnum() MockApiBaseInfoStatusCodeEnum {
	return MockApiBaseInfoStatusCodeEnum{
		E_200: MockApiBaseInfoStatusCode{
			value: 200,
		}, E_201: MockApiBaseInfoStatusCode{
			value: 201,
		}, E_202: MockApiBaseInfoStatusCode{
			value: 202,
		}, E_203: MockApiBaseInfoStatusCode{
			value: 203,
		}, E_204: MockApiBaseInfoStatusCode{
			value: 204,
		}, E_205: MockApiBaseInfoStatusCode{
			value: 205,
		}, E_206: MockApiBaseInfoStatusCode{
			value: 206,
		}, E_300: MockApiBaseInfoStatusCode{
			value: 300,
		}, E_301: MockApiBaseInfoStatusCode{
			value: 301,
		}, E_302: MockApiBaseInfoStatusCode{
			value: 302,
		}, E_303: MockApiBaseInfoStatusCode{
			value: 303,
		}, E_304: MockApiBaseInfoStatusCode{
			value: 304,
		}, E_305: MockApiBaseInfoStatusCode{
			value: 305,
		}, E_306: MockApiBaseInfoStatusCode{
			value: 306,
		}, E_307: MockApiBaseInfoStatusCode{
			value: 307,
		}, E_400: MockApiBaseInfoStatusCode{
			value: 400,
		}, E_401: MockApiBaseInfoStatusCode{
			value: 401,
		}, E_402: MockApiBaseInfoStatusCode{
			value: 402,
		}, E_403: MockApiBaseInfoStatusCode{
			value: 403,
		}, E_404: MockApiBaseInfoStatusCode{
			value: 404,
		}, E_405: MockApiBaseInfoStatusCode{
			value: 405,
		}, E_406: MockApiBaseInfoStatusCode{
			value: 406,
		}, E_407: MockApiBaseInfoStatusCode{
			value: 407,
		}, E_408: MockApiBaseInfoStatusCode{
			value: 408,
		}, E_409: MockApiBaseInfoStatusCode{
			value: 409,
		}, E_410: MockApiBaseInfoStatusCode{
			value: 410,
		}, E_411: MockApiBaseInfoStatusCode{
			value: 411,
		}, E_412: MockApiBaseInfoStatusCode{
			value: 412,
		}, E_413: MockApiBaseInfoStatusCode{
			value: 413,
		}, E_414: MockApiBaseInfoStatusCode{
			value: 414,
		}, E_415: MockApiBaseInfoStatusCode{
			value: 415,
		}, E_416: MockApiBaseInfoStatusCode{
			value: 416,
		}, E_417: MockApiBaseInfoStatusCode{
			value: 417,
		}, E_450: MockApiBaseInfoStatusCode{
			value: 450,
		}, E_451: MockApiBaseInfoStatusCode{
			value: 451,
		}, E_500: MockApiBaseInfoStatusCode{
			value: 500,
		}, E_501: MockApiBaseInfoStatusCode{
			value: 501,
		}, E_502: MockApiBaseInfoStatusCode{
			value: 502,
		}, E_503: MockApiBaseInfoStatusCode{
			value: 503,
		}, E_504: MockApiBaseInfoStatusCode{
			value: 504,
		}, E_505: MockApiBaseInfoStatusCode{
			value: 505,
		},
	}
}

func (c MockApiBaseInfoStatusCode) Value() int32 {
	return c.value
}

func (c MockApiBaseInfoStatusCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MockApiBaseInfoStatusCode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
