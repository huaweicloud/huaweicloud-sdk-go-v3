package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyMockResp struct {
	// 编号

	Id *string `json:"id,omitempty"`
	// 策略后端名称。字符串由中文、英文字母、数字、下划线组成，且只能以中文或英文开头。

	Name string `json:"name"`
	// 策略条件列表

	Conditions []ConditionResp `json:"conditions"`
	// 后端参数列表

	BackendParams *[]BackendParam `json:"backend_params,omitempty"`
	// 关联的策略组合模式： - ALL：满足全部条件 - ANY：满足任一条件

	EffectMode ApiPolicyMockRespEffectMode `json:"effect_mode"`
	// 后端自定义认证对象的ID

	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// 返回结果

	ResultContent *string `json:"result_content,omitempty"`
	// mock后端自定义状态码： \"200\": \"OK\", \"201\": \"Created\", \"202\": \"Accepted\", \"203\": \"NonAuthoritativeInformation\", \"204\": \"NoContent\", \"205\": \"ResetContent\", \"206\": \"PartialContent\", \"300\": \"MultipleChoices\", \"301\": \"MovedPermanently\", \"302\": \"Found\", \"303\": \"SeeOther\", \"304\": \"NotModified\", \"305\": \"UseProxy\", \"306\": \"Unused\", \"307\": \"TemporaryRedirect\", \"400\": \"BadRequest\", \"401\": \"Unauthorized\", \"402\": \"PaymentRequired\", \"403\": \"Forbidden\", \"404\": \"NotFound\", \"405\": \"MethodNotAllowed\", \"406\": \"NotAcceptable\", \"407\": \"ProxyAuthenticationRequired\", \"408\": \"RequestTimeout\", \"409\": \"Conflict\", \"410\": \"Gone\", \"411\": \"LengthRequired\", \"412\": \"PreconditionFailed\", \"413\": \"RequestEntityTooLarge\", \"414\": \"RequestURITooLong\", \"415\": \"UnsupportedMediaType\", \"416\": \"RequestedRangeNotSatisfiable\", \"417\": \"ExpectationFailed\", \"450\": \"ParameterRequried\", \"451\": \"MethodConnectException\", \"500\": \"InternalServerError\", \"501\": \"NotImplemented\", \"502\": \"BadGateway\", \"503\": \"ServiceUnavailable\", \"504\": \"GatewayTimeout\", \"505\": \"HTTPVersionNotSupported\",

	StatusCode *ApiPolicyMockRespStatusCode `json:"status_code,omitempty"`
	// mock后端自定义响应头header

	Header *[]MockApiBaseInfoHeader `json:"header,omitempty"`
}

func (o ApiPolicyMockResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyMockResp struct{}"
	}

	return strings.Join([]string{"ApiPolicyMockResp", string(data)}, " ")
}

type ApiPolicyMockRespEffectMode struct {
	value string
}

type ApiPolicyMockRespEffectModeEnum struct {
	ALL ApiPolicyMockRespEffectMode
	ANY ApiPolicyMockRespEffectMode
}

func GetApiPolicyMockRespEffectModeEnum() ApiPolicyMockRespEffectModeEnum {
	return ApiPolicyMockRespEffectModeEnum{
		ALL: ApiPolicyMockRespEffectMode{
			value: "ALL",
		},
		ANY: ApiPolicyMockRespEffectMode{
			value: "ANY",
		},
	}
}

func (c ApiPolicyMockRespEffectMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyMockRespEffectMode) UnmarshalJSON(b []byte) error {
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

type ApiPolicyMockRespStatusCode struct {
	value int32
}

type ApiPolicyMockRespStatusCodeEnum struct {
	E_200 ApiPolicyMockRespStatusCode
	E_201 ApiPolicyMockRespStatusCode
	E_202 ApiPolicyMockRespStatusCode
	E_203 ApiPolicyMockRespStatusCode
	E_204 ApiPolicyMockRespStatusCode
	E_205 ApiPolicyMockRespStatusCode
	E_206 ApiPolicyMockRespStatusCode
	E_300 ApiPolicyMockRespStatusCode
	E_301 ApiPolicyMockRespStatusCode
	E_302 ApiPolicyMockRespStatusCode
	E_303 ApiPolicyMockRespStatusCode
	E_304 ApiPolicyMockRespStatusCode
	E_305 ApiPolicyMockRespStatusCode
	E_306 ApiPolicyMockRespStatusCode
	E_307 ApiPolicyMockRespStatusCode
	E_400 ApiPolicyMockRespStatusCode
	E_401 ApiPolicyMockRespStatusCode
	E_402 ApiPolicyMockRespStatusCode
	E_403 ApiPolicyMockRespStatusCode
	E_404 ApiPolicyMockRespStatusCode
	E_405 ApiPolicyMockRespStatusCode
	E_406 ApiPolicyMockRespStatusCode
	E_407 ApiPolicyMockRespStatusCode
	E_408 ApiPolicyMockRespStatusCode
	E_409 ApiPolicyMockRespStatusCode
	E_410 ApiPolicyMockRespStatusCode
	E_411 ApiPolicyMockRespStatusCode
	E_412 ApiPolicyMockRespStatusCode
	E_413 ApiPolicyMockRespStatusCode
	E_414 ApiPolicyMockRespStatusCode
	E_415 ApiPolicyMockRespStatusCode
	E_416 ApiPolicyMockRespStatusCode
	E_417 ApiPolicyMockRespStatusCode
	E_450 ApiPolicyMockRespStatusCode
	E_451 ApiPolicyMockRespStatusCode
	E_500 ApiPolicyMockRespStatusCode
	E_501 ApiPolicyMockRespStatusCode
	E_502 ApiPolicyMockRespStatusCode
	E_503 ApiPolicyMockRespStatusCode
	E_504 ApiPolicyMockRespStatusCode
	E_505 ApiPolicyMockRespStatusCode
}

func GetApiPolicyMockRespStatusCodeEnum() ApiPolicyMockRespStatusCodeEnum {
	return ApiPolicyMockRespStatusCodeEnum{
		E_200: ApiPolicyMockRespStatusCode{
			value: 200,
		}, E_201: ApiPolicyMockRespStatusCode{
			value: 201,
		}, E_202: ApiPolicyMockRespStatusCode{
			value: 202,
		}, E_203: ApiPolicyMockRespStatusCode{
			value: 203,
		}, E_204: ApiPolicyMockRespStatusCode{
			value: 204,
		}, E_205: ApiPolicyMockRespStatusCode{
			value: 205,
		}, E_206: ApiPolicyMockRespStatusCode{
			value: 206,
		}, E_300: ApiPolicyMockRespStatusCode{
			value: 300,
		}, E_301: ApiPolicyMockRespStatusCode{
			value: 301,
		}, E_302: ApiPolicyMockRespStatusCode{
			value: 302,
		}, E_303: ApiPolicyMockRespStatusCode{
			value: 303,
		}, E_304: ApiPolicyMockRespStatusCode{
			value: 304,
		}, E_305: ApiPolicyMockRespStatusCode{
			value: 305,
		}, E_306: ApiPolicyMockRespStatusCode{
			value: 306,
		}, E_307: ApiPolicyMockRespStatusCode{
			value: 307,
		}, E_400: ApiPolicyMockRespStatusCode{
			value: 400,
		}, E_401: ApiPolicyMockRespStatusCode{
			value: 401,
		}, E_402: ApiPolicyMockRespStatusCode{
			value: 402,
		}, E_403: ApiPolicyMockRespStatusCode{
			value: 403,
		}, E_404: ApiPolicyMockRespStatusCode{
			value: 404,
		}, E_405: ApiPolicyMockRespStatusCode{
			value: 405,
		}, E_406: ApiPolicyMockRespStatusCode{
			value: 406,
		}, E_407: ApiPolicyMockRespStatusCode{
			value: 407,
		}, E_408: ApiPolicyMockRespStatusCode{
			value: 408,
		}, E_409: ApiPolicyMockRespStatusCode{
			value: 409,
		}, E_410: ApiPolicyMockRespStatusCode{
			value: 410,
		}, E_411: ApiPolicyMockRespStatusCode{
			value: 411,
		}, E_412: ApiPolicyMockRespStatusCode{
			value: 412,
		}, E_413: ApiPolicyMockRespStatusCode{
			value: 413,
		}, E_414: ApiPolicyMockRespStatusCode{
			value: 414,
		}, E_415: ApiPolicyMockRespStatusCode{
			value: 415,
		}, E_416: ApiPolicyMockRespStatusCode{
			value: 416,
		}, E_417: ApiPolicyMockRespStatusCode{
			value: 417,
		}, E_450: ApiPolicyMockRespStatusCode{
			value: 450,
		}, E_451: ApiPolicyMockRespStatusCode{
			value: 451,
		}, E_500: ApiPolicyMockRespStatusCode{
			value: 500,
		}, E_501: ApiPolicyMockRespStatusCode{
			value: 501,
		}, E_502: ApiPolicyMockRespStatusCode{
			value: 502,
		}, E_503: ApiPolicyMockRespStatusCode{
			value: 503,
		}, E_504: ApiPolicyMockRespStatusCode{
			value: 504,
		}, E_505: ApiPolicyMockRespStatusCode{
			value: 505,
		},
	}
}

func (c ApiPolicyMockRespStatusCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyMockRespStatusCode) UnmarshalJSON(b []byte) error {
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
