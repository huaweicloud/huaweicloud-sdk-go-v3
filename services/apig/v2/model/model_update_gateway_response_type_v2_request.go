package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type UpdateGatewayResponseTypeV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 分组的编号

	GroupId string `json:"group_id"`
	// 响应编号

	ResponseId string `json:"response_id"`
	// 错误类型

	ResponseType UpdateGatewayResponseTypeV2RequestResponseType `json:"response_type"`

	Body *ResponseInfo `json:"body,omitempty"`
}

func (o UpdateGatewayResponseTypeV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGatewayResponseTypeV2Request struct{}"
	}

	return strings.Join([]string{"UpdateGatewayResponseTypeV2Request", string(data)}, " ")
}

type UpdateGatewayResponseTypeV2RequestResponseType struct {
	value string
}

type UpdateGatewayResponseTypeV2RequestResponseTypeEnum struct {
	AUTH_FAILURE                  UpdateGatewayResponseTypeV2RequestResponseType
	AUTH_HEADER_MISSING           UpdateGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_FAILURE            UpdateGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_CONF_FAILURE       UpdateGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_IDENTITIES_FAILURE UpdateGatewayResponseTypeV2RequestResponseType
	BACKEND_UNAVAILABLE           UpdateGatewayResponseTypeV2RequestResponseType
	BACKEND_TIMEOUT               UpdateGatewayResponseTypeV2RequestResponseType
	THROTTLED                     UpdateGatewayResponseTypeV2RequestResponseType
	UNAUTHORIZED                  UpdateGatewayResponseTypeV2RequestResponseType
	ACCESS_DENIED                 UpdateGatewayResponseTypeV2RequestResponseType
	NOT_FOUND                     UpdateGatewayResponseTypeV2RequestResponseType
	REQUEST_PARAMETERS_FAILURE    UpdateGatewayResponseTypeV2RequestResponseType
	DEFAULT_4_XX                  UpdateGatewayResponseTypeV2RequestResponseType
	DEFAULT_5_XX                  UpdateGatewayResponseTypeV2RequestResponseType
}

func GetUpdateGatewayResponseTypeV2RequestResponseTypeEnum() UpdateGatewayResponseTypeV2RequestResponseTypeEnum {
	return UpdateGatewayResponseTypeV2RequestResponseTypeEnum{
		AUTH_FAILURE: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "AUTH_FAILURE",
		},
		AUTH_HEADER_MISSING: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "AUTH_HEADER_MISSING",
		},
		AUTHORIZER_FAILURE: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_FAILURE",
		},
		AUTHORIZER_CONF_FAILURE: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_CONF_FAILURE",
		},
		AUTHORIZER_IDENTITIES_FAILURE: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_IDENTITIES_FAILURE",
		},
		BACKEND_UNAVAILABLE: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "BACKEND_UNAVAILABLE",
		},
		BACKEND_TIMEOUT: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "BACKEND_TIMEOUT",
		},
		THROTTLED: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "THROTTLED",
		},
		UNAUTHORIZED: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "UNAUTHORIZED",
		},
		ACCESS_DENIED: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "ACCESS_DENIED",
		},
		NOT_FOUND: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "NOT_FOUND",
		},
		REQUEST_PARAMETERS_FAILURE: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "REQUEST_PARAMETERS_FAILURE",
		},
		DEFAULT_4_XX: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "DEFAULT_4XX",
		},
		DEFAULT_5_XX: UpdateGatewayResponseTypeV2RequestResponseType{
			value: "DEFAULT_5XX",
		},
	}
}

func (c UpdateGatewayResponseTypeV2RequestResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateGatewayResponseTypeV2RequestResponseType) UnmarshalJSON(b []byte) error {
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
