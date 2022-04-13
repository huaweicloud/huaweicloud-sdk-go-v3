package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ApiPolicyMockCreate struct {
	// 返回结果

	ResultContent *string `json:"result_content,omitempty"`
	// mock后端自定义状态码： \"200\": \"OK\", \"201\": \"Created\", \"202\": \"Accepted\", \"203\": \"NonAuthoritativeInformation\", \"204\": \"NoContent\", \"205\": \"ResetContent\", \"206\": \"PartialContent\", \"300\": \"MultipleChoices\", \"301\": \"MovedPermanently\", \"302\": \"Found\", \"303\": \"SeeOther\", \"304\": \"NotModified\", \"305\": \"UseProxy\", \"306\": \"Unused\", \"307\": \"TemporaryRedirect\", \"400\": \"BadRequest\", \"401\": \"Unauthorized\", \"402\": \"PaymentRequired\", \"403\": \"Forbidden\", \"404\": \"NotFound\", \"405\": \"MethodNotAllowed\", \"406\": \"NotAcceptable\", \"407\": \"ProxyAuthenticationRequired\", \"408\": \"RequestTimeout\", \"409\": \"Conflict\", \"410\": \"Gone\", \"411\": \"LengthRequired\", \"412\": \"PreconditionFailed\", \"413\": \"RequestEntityTooLarge\", \"414\": \"RequestURITooLong\", \"415\": \"UnsupportedMediaType\", \"416\": \"RequestedRangeNotSatisfiable\", \"417\": \"ExpectationFailed\", \"450\": \"ParameterRequried\", \"451\": \"MethodConnectException\", \"500\": \"InternalServerError\", \"501\": \"NotImplemented\", \"502\": \"BadGateway\", \"503\": \"ServiceUnavailable\", \"504\": \"GatewayTimeout\", \"505\": \"HTTPVersionNotSupported\",

	StatusCode *ApiPolicyMockCreateStatusCode `json:"status_code,omitempty"`
	// mock后端自定义响应头header

	Header *[]MockApiBaseInfoHeader `json:"header,omitempty"`
	// 关联的策略组合模式： - ALL：满足全部条件 - ANY：满足任一条件

	EffectMode ApiPolicyMockCreateEffectMode `json:"effect_mode"`
	// 策略后端名称。字符串由中文、英文字母、数字、下划线组成，且只能以中文或英文开头。

	Name string `json:"name"`
	// 后端参数列表

	BackendParams *[]BackendParamBase `json:"backend_params,omitempty"`
	// 策略条件列表

	Conditions []ApiConditionCreate `json:"conditions"`
	// 后端自定义认证对象的ID

	AuthorizerId *string `json:"authorizer_id,omitempty"`
}

func (o ApiPolicyMockCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyMockCreate struct{}"
	}

	return strings.Join([]string{"ApiPolicyMockCreate", string(data)}, " ")
}

type ApiPolicyMockCreateStatusCode struct {
	value int32
}

type ApiPolicyMockCreateStatusCodeEnum struct {
	E_200 ApiPolicyMockCreateStatusCode
	E_201 ApiPolicyMockCreateStatusCode
	E_202 ApiPolicyMockCreateStatusCode
	E_203 ApiPolicyMockCreateStatusCode
	E_204 ApiPolicyMockCreateStatusCode
	E_205 ApiPolicyMockCreateStatusCode
	E_206 ApiPolicyMockCreateStatusCode
	E_300 ApiPolicyMockCreateStatusCode
	E_301 ApiPolicyMockCreateStatusCode
	E_302 ApiPolicyMockCreateStatusCode
	E_303 ApiPolicyMockCreateStatusCode
	E_304 ApiPolicyMockCreateStatusCode
	E_305 ApiPolicyMockCreateStatusCode
	E_306 ApiPolicyMockCreateStatusCode
	E_307 ApiPolicyMockCreateStatusCode
	E_400 ApiPolicyMockCreateStatusCode
	E_401 ApiPolicyMockCreateStatusCode
	E_402 ApiPolicyMockCreateStatusCode
	E_403 ApiPolicyMockCreateStatusCode
	E_404 ApiPolicyMockCreateStatusCode
	E_405 ApiPolicyMockCreateStatusCode
	E_406 ApiPolicyMockCreateStatusCode
	E_407 ApiPolicyMockCreateStatusCode
	E_408 ApiPolicyMockCreateStatusCode
	E_409 ApiPolicyMockCreateStatusCode
	E_410 ApiPolicyMockCreateStatusCode
	E_411 ApiPolicyMockCreateStatusCode
	E_412 ApiPolicyMockCreateStatusCode
	E_413 ApiPolicyMockCreateStatusCode
	E_414 ApiPolicyMockCreateStatusCode
	E_415 ApiPolicyMockCreateStatusCode
	E_416 ApiPolicyMockCreateStatusCode
	E_417 ApiPolicyMockCreateStatusCode
	E_450 ApiPolicyMockCreateStatusCode
	E_451 ApiPolicyMockCreateStatusCode
	E_500 ApiPolicyMockCreateStatusCode
	E_501 ApiPolicyMockCreateStatusCode
	E_502 ApiPolicyMockCreateStatusCode
	E_503 ApiPolicyMockCreateStatusCode
	E_504 ApiPolicyMockCreateStatusCode
	E_505 ApiPolicyMockCreateStatusCode
}

func GetApiPolicyMockCreateStatusCodeEnum() ApiPolicyMockCreateStatusCodeEnum {
	return ApiPolicyMockCreateStatusCodeEnum{
		E_200: ApiPolicyMockCreateStatusCode{
			value: 200,
		}, E_201: ApiPolicyMockCreateStatusCode{
			value: 201,
		}, E_202: ApiPolicyMockCreateStatusCode{
			value: 202,
		}, E_203: ApiPolicyMockCreateStatusCode{
			value: 203,
		}, E_204: ApiPolicyMockCreateStatusCode{
			value: 204,
		}, E_205: ApiPolicyMockCreateStatusCode{
			value: 205,
		}, E_206: ApiPolicyMockCreateStatusCode{
			value: 206,
		}, E_300: ApiPolicyMockCreateStatusCode{
			value: 300,
		}, E_301: ApiPolicyMockCreateStatusCode{
			value: 301,
		}, E_302: ApiPolicyMockCreateStatusCode{
			value: 302,
		}, E_303: ApiPolicyMockCreateStatusCode{
			value: 303,
		}, E_304: ApiPolicyMockCreateStatusCode{
			value: 304,
		}, E_305: ApiPolicyMockCreateStatusCode{
			value: 305,
		}, E_306: ApiPolicyMockCreateStatusCode{
			value: 306,
		}, E_307: ApiPolicyMockCreateStatusCode{
			value: 307,
		}, E_400: ApiPolicyMockCreateStatusCode{
			value: 400,
		}, E_401: ApiPolicyMockCreateStatusCode{
			value: 401,
		}, E_402: ApiPolicyMockCreateStatusCode{
			value: 402,
		}, E_403: ApiPolicyMockCreateStatusCode{
			value: 403,
		}, E_404: ApiPolicyMockCreateStatusCode{
			value: 404,
		}, E_405: ApiPolicyMockCreateStatusCode{
			value: 405,
		}, E_406: ApiPolicyMockCreateStatusCode{
			value: 406,
		}, E_407: ApiPolicyMockCreateStatusCode{
			value: 407,
		}, E_408: ApiPolicyMockCreateStatusCode{
			value: 408,
		}, E_409: ApiPolicyMockCreateStatusCode{
			value: 409,
		}, E_410: ApiPolicyMockCreateStatusCode{
			value: 410,
		}, E_411: ApiPolicyMockCreateStatusCode{
			value: 411,
		}, E_412: ApiPolicyMockCreateStatusCode{
			value: 412,
		}, E_413: ApiPolicyMockCreateStatusCode{
			value: 413,
		}, E_414: ApiPolicyMockCreateStatusCode{
			value: 414,
		}, E_415: ApiPolicyMockCreateStatusCode{
			value: 415,
		}, E_416: ApiPolicyMockCreateStatusCode{
			value: 416,
		}, E_417: ApiPolicyMockCreateStatusCode{
			value: 417,
		}, E_450: ApiPolicyMockCreateStatusCode{
			value: 450,
		}, E_451: ApiPolicyMockCreateStatusCode{
			value: 451,
		}, E_500: ApiPolicyMockCreateStatusCode{
			value: 500,
		}, E_501: ApiPolicyMockCreateStatusCode{
			value: 501,
		}, E_502: ApiPolicyMockCreateStatusCode{
			value: 502,
		}, E_503: ApiPolicyMockCreateStatusCode{
			value: 503,
		}, E_504: ApiPolicyMockCreateStatusCode{
			value: 504,
		}, E_505: ApiPolicyMockCreateStatusCode{
			value: 505,
		},
	}
}

func (c ApiPolicyMockCreateStatusCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyMockCreateStatusCode) UnmarshalJSON(b []byte) error {
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

type ApiPolicyMockCreateEffectMode struct {
	value string
}

type ApiPolicyMockCreateEffectModeEnum struct {
	ALL ApiPolicyMockCreateEffectMode
	ANY ApiPolicyMockCreateEffectMode
}

func GetApiPolicyMockCreateEffectModeEnum() ApiPolicyMockCreateEffectModeEnum {
	return ApiPolicyMockCreateEffectModeEnum{
		ALL: ApiPolicyMockCreateEffectMode{
			value: "ALL",
		},
		ANY: ApiPolicyMockCreateEffectMode{
			value: "ANY",
		},
	}
}

func (c ApiPolicyMockCreateEffectMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyMockCreateEffectMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
