package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDetailsOfGatewayResponseTypeV2Request Request Object
type ShowDetailsOfGatewayResponseTypeV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	// 响应编号
	ResponseId string `json:"response_id"`

	// 错误类型 - AUTH_FAILURE: 认证失败，IAM或APP认证校验失败 - AUTH_HEADER_MISSING: 认证身份来源信息缺失 - AUTHORIZER_FAILURE: 自定义认证方返回认证失败 - AUTHORIZER_CONF_FAILURE:自定义认证方异常，通信失败、返回异常响应等错误 - AUTHORIZER_IDENTITIES_FAILURE: 前端自定义认证的身份来源信息缺失或不合法错误 - BACKEND_UNAVAILABLE: 后端不可用，网络不可达错误 - BACKEND_TIMEOUT: 后端超时，与后端的网络交互超过预配置的时间错误 - THROTTLED: API调用次数超出所配置的流量策略阈值 - UNAUTHORIZED: 使用的凭据未被授权访问该API - ACCESS_DENIED: 拒绝访问，如触发配置的访问控制策略、或异常攻击检测拦截 - NOT_FOUND: 未匹配到API错误 - REQUEST_PARAMETERS_FAILURE: 请求参数校验失败、不支持的HTTP方法 - DEFAULT_4XX: 其它4XX类错误 - DEFAULT_5XX: 其它5XX类错误 - THIRD_AUTH_FAILURE: 第三方认证方返回认证失败 - THIRD_AUTH_IDENTITIES_FAILURE: 第三方认证的身份来源信息缺失或不合法错误 - THIRD_AUTH_CONF_FAILURE: 第三方认证方异常，通信失败、返回异常响应等错误 - ORCHESTRATION_PARAMETER_NOT_FOUND: 参数编排失败，请求中没有待编排的入参 - ORCHESTRATION_FAILURE: 参数编排失败，没有编排规则匹配成功
	ResponseType ShowDetailsOfGatewayResponseTypeV2RequestResponseType `json:"response_type"`
}

func (o ShowDetailsOfGatewayResponseTypeV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfGatewayResponseTypeV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfGatewayResponseTypeV2Request", string(data)}, " ")
}

type ShowDetailsOfGatewayResponseTypeV2RequestResponseType struct {
	value string
}

type ShowDetailsOfGatewayResponseTypeV2RequestResponseTypeEnum struct {
	AUTH_FAILURE                      ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	AUTH_HEADER_MISSING               ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_FAILURE                ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_CONF_FAILURE           ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_IDENTITIES_FAILURE     ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	BACKEND_UNAVAILABLE               ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	BACKEND_TIMEOUT                   ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	THROTTLED                         ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	UNAUTHORIZED                      ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	ACCESS_DENIED                     ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	NOT_FOUND                         ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	REQUEST_PARAMETERS_FAILURE        ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	DEFAULT_4_XX                      ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	DEFAULT_5_XX                      ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	THIRD_AUTH_FAILURE                ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	THIRD_AUTH_IDENTITIES_FAILURE     ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	THIRD_AUTH_CONF_FAILURE           ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	ORCHESTRATION_PARAMETER_NOT_FOUND ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	ORCHESTRATION_FAILURE             ShowDetailsOfGatewayResponseTypeV2RequestResponseType
}

func GetShowDetailsOfGatewayResponseTypeV2RequestResponseTypeEnum() ShowDetailsOfGatewayResponseTypeV2RequestResponseTypeEnum {
	return ShowDetailsOfGatewayResponseTypeV2RequestResponseTypeEnum{
		AUTH_FAILURE: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "AUTH_FAILURE",
		},
		AUTH_HEADER_MISSING: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "AUTH_HEADER_MISSING",
		},
		AUTHORIZER_FAILURE: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_FAILURE",
		},
		AUTHORIZER_CONF_FAILURE: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_CONF_FAILURE",
		},
		AUTHORIZER_IDENTITIES_FAILURE: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_IDENTITIES_FAILURE",
		},
		BACKEND_UNAVAILABLE: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "BACKEND_UNAVAILABLE",
		},
		BACKEND_TIMEOUT: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "BACKEND_TIMEOUT",
		},
		THROTTLED: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "THROTTLED",
		},
		UNAUTHORIZED: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "UNAUTHORIZED",
		},
		ACCESS_DENIED: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "ACCESS_DENIED",
		},
		NOT_FOUND: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "NOT_FOUND",
		},
		REQUEST_PARAMETERS_FAILURE: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "REQUEST_PARAMETERS_FAILURE",
		},
		DEFAULT_4_XX: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "DEFAULT_4XX",
		},
		DEFAULT_5_XX: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "DEFAULT_5XX",
		},
		THIRD_AUTH_FAILURE: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "THIRD_AUTH_FAILURE",
		},
		THIRD_AUTH_IDENTITIES_FAILURE: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "THIRD_AUTH_IDENTITIES_FAILURE",
		},
		THIRD_AUTH_CONF_FAILURE: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "THIRD_AUTH_CONF_FAILURE",
		},
		ORCHESTRATION_PARAMETER_NOT_FOUND: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "ORCHESTRATION_PARAMETER_NOT_FOUND",
		},
		ORCHESTRATION_FAILURE: ShowDetailsOfGatewayResponseTypeV2RequestResponseType{
			value: "ORCHESTRATION_FAILURE",
		},
	}
}

func (c ShowDetailsOfGatewayResponseTypeV2RequestResponseType) Value() string {
	return c.value
}

func (c ShowDetailsOfGatewayResponseTypeV2RequestResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfGatewayResponseTypeV2RequestResponseType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
