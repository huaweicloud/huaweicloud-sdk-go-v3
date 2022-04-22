package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// mock后端详情
type ApiMock struct {

	// 描述信息。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 返回结果
	ResultContent *string `json:"result_content,omitempty"`

	// 版本。字符长度不超过64
	Version *string `json:"version,omitempty"`

	// 后端自定义认证ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// mock后端自定义状态码： \"200\": \"OK\", \"201\": \"Created\", \"202\": \"Accepted\", \"203\": \"NonAuthoritativeInformation\", \"204\": \"NoContent\", \"205\": \"ResetContent\", \"206\": \"PartialContent\", \"300\": \"MultipleChoices\", \"301\": \"MovedPermanently\", \"302\": \"Found\", \"303\": \"SeeOther\", \"304\": \"NotModified\", \"305\": \"UseProxy\", \"306\": \"Unused\", \"307\": \"TemporaryRedirect\", \"400\": \"BadRequest\", \"401\": \"Unauthorized\", \"402\": \"PaymentRequired\", \"403\": \"Forbidden\", \"404\": \"NotFound\", \"405\": \"MethodNotAllowed\", \"406\": \"NotAcceptable\", \"407\": \"ProxyAuthenticationRequired\", \"408\": \"RequestTimeout\", \"409\": \"Conflict\", \"410\": \"Gone\", \"411\": \"LengthRequired\", \"412\": \"PreconditionFailed\", \"413\": \"RequestEntityTooLarge\", \"414\": \"RequestURITooLong\", \"415\": \"UnsupportedMediaType\", \"416\": \"RequestedRangeNotSatisfiable\", \"417\": \"ExpectationFailed\", \"450\": \"ParameterRequried\", \"451\": \"MethodConnectException\", \"500\": \"InternalServerError\", \"501\": \"NotImplemented\", \"502\": \"BadGateway\", \"503\": \"ServiceUnavailable\", \"504\": \"GatewayTimeout\", \"505\": \"HTTPVersionNotSupported\",
	StatusCode *ApiMockStatusCode `json:"status_code,omitempty"`

	// mock后端自定义响应头header
	Header *[]MockApiBaseInfoHeader `json:"header,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// 后端状态   - 1： 有效
	Status *ApiMockStatus `json:"status,omitempty"`

	// 修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o ApiMock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiMock struct{}"
	}

	return strings.Join([]string{"ApiMock", string(data)}, " ")
}

type ApiMockStatusCode struct {
	value int32
}

type ApiMockStatusCodeEnum struct {
	E_200 ApiMockStatusCode
	E_201 ApiMockStatusCode
	E_202 ApiMockStatusCode
	E_203 ApiMockStatusCode
	E_204 ApiMockStatusCode
	E_205 ApiMockStatusCode
	E_206 ApiMockStatusCode
	E_300 ApiMockStatusCode
	E_301 ApiMockStatusCode
	E_302 ApiMockStatusCode
	E_303 ApiMockStatusCode
	E_304 ApiMockStatusCode
	E_305 ApiMockStatusCode
	E_306 ApiMockStatusCode
	E_307 ApiMockStatusCode
	E_400 ApiMockStatusCode
	E_401 ApiMockStatusCode
	E_402 ApiMockStatusCode
	E_403 ApiMockStatusCode
	E_404 ApiMockStatusCode
	E_405 ApiMockStatusCode
	E_406 ApiMockStatusCode
	E_407 ApiMockStatusCode
	E_408 ApiMockStatusCode
	E_409 ApiMockStatusCode
	E_410 ApiMockStatusCode
	E_411 ApiMockStatusCode
	E_412 ApiMockStatusCode
	E_413 ApiMockStatusCode
	E_414 ApiMockStatusCode
	E_415 ApiMockStatusCode
	E_416 ApiMockStatusCode
	E_417 ApiMockStatusCode
	E_450 ApiMockStatusCode
	E_451 ApiMockStatusCode
	E_500 ApiMockStatusCode
	E_501 ApiMockStatusCode
	E_502 ApiMockStatusCode
	E_503 ApiMockStatusCode
	E_504 ApiMockStatusCode
	E_505 ApiMockStatusCode
}

func GetApiMockStatusCodeEnum() ApiMockStatusCodeEnum {
	return ApiMockStatusCodeEnum{
		E_200: ApiMockStatusCode{
			value: 200,
		}, E_201: ApiMockStatusCode{
			value: 201,
		}, E_202: ApiMockStatusCode{
			value: 202,
		}, E_203: ApiMockStatusCode{
			value: 203,
		}, E_204: ApiMockStatusCode{
			value: 204,
		}, E_205: ApiMockStatusCode{
			value: 205,
		}, E_206: ApiMockStatusCode{
			value: 206,
		}, E_300: ApiMockStatusCode{
			value: 300,
		}, E_301: ApiMockStatusCode{
			value: 301,
		}, E_302: ApiMockStatusCode{
			value: 302,
		}, E_303: ApiMockStatusCode{
			value: 303,
		}, E_304: ApiMockStatusCode{
			value: 304,
		}, E_305: ApiMockStatusCode{
			value: 305,
		}, E_306: ApiMockStatusCode{
			value: 306,
		}, E_307: ApiMockStatusCode{
			value: 307,
		}, E_400: ApiMockStatusCode{
			value: 400,
		}, E_401: ApiMockStatusCode{
			value: 401,
		}, E_402: ApiMockStatusCode{
			value: 402,
		}, E_403: ApiMockStatusCode{
			value: 403,
		}, E_404: ApiMockStatusCode{
			value: 404,
		}, E_405: ApiMockStatusCode{
			value: 405,
		}, E_406: ApiMockStatusCode{
			value: 406,
		}, E_407: ApiMockStatusCode{
			value: 407,
		}, E_408: ApiMockStatusCode{
			value: 408,
		}, E_409: ApiMockStatusCode{
			value: 409,
		}, E_410: ApiMockStatusCode{
			value: 410,
		}, E_411: ApiMockStatusCode{
			value: 411,
		}, E_412: ApiMockStatusCode{
			value: 412,
		}, E_413: ApiMockStatusCode{
			value: 413,
		}, E_414: ApiMockStatusCode{
			value: 414,
		}, E_415: ApiMockStatusCode{
			value: 415,
		}, E_416: ApiMockStatusCode{
			value: 416,
		}, E_417: ApiMockStatusCode{
			value: 417,
		}, E_450: ApiMockStatusCode{
			value: 450,
		}, E_451: ApiMockStatusCode{
			value: 451,
		}, E_500: ApiMockStatusCode{
			value: 500,
		}, E_501: ApiMockStatusCode{
			value: 501,
		}, E_502: ApiMockStatusCode{
			value: 502,
		}, E_503: ApiMockStatusCode{
			value: 503,
		}, E_504: ApiMockStatusCode{
			value: 504,
		}, E_505: ApiMockStatusCode{
			value: 505,
		},
	}
}

func (c ApiMockStatusCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiMockStatusCode) UnmarshalJSON(b []byte) error {
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

type ApiMockStatus struct {
	value int32
}

type ApiMockStatusEnum struct {
	E_1 ApiMockStatus
}

func GetApiMockStatusEnum() ApiMockStatusEnum {
	return ApiMockStatusEnum{
		E_1: ApiMockStatus{
			value: 1,
		},
	}
}

func (c ApiMockStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiMockStatus) UnmarshalJSON(b []byte) error {
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
