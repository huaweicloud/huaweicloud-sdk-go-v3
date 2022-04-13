package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ShowDetailsOfGatewayResponseTypeV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 分组的编号

	GroupId string `json:"group_id"`
	// 响应编号

	ResponseId string `json:"response_id"`
	// 错误类型

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
	AUTH_FAILURE                  ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	AUTH_HEADER_MISSING           ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_FAILURE            ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_CONF_FAILURE       ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_IDENTITIES_FAILURE ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	BACKEND_UNAVAILABLE           ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	BACKEND_TIMEOUT               ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	THROTTLED                     ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	UNAUTHORIZED                  ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	ACCESS_DENIED                 ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	NOT_FOUND                     ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	REQUEST_PARAMETERS_FAILURE    ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	DEFAULT_4_XX                  ShowDetailsOfGatewayResponseTypeV2RequestResponseType
	DEFAULT_5_XX                  ShowDetailsOfGatewayResponseTypeV2RequestResponseType
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
	}
}

func (c ShowDetailsOfGatewayResponseTypeV2RequestResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfGatewayResponseTypeV2RequestResponseType) UnmarshalJSON(b []byte) error {
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
