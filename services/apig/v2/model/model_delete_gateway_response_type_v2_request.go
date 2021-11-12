package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteGatewayResponseTypeV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 分组的编号

	GroupId string `json:"group_id"`
	// 响应编号

	ResponseId string `json:"response_id"`
	// 错误类型

	ResponseType DeleteGatewayResponseTypeV2RequestResponseType `json:"response_type"`
}

func (o DeleteGatewayResponseTypeV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGatewayResponseTypeV2Request struct{}"
	}

	return strings.Join([]string{"DeleteGatewayResponseTypeV2Request", string(data)}, " ")
}

type DeleteGatewayResponseTypeV2RequestResponseType struct {
	value string
}

type DeleteGatewayResponseTypeV2RequestResponseTypeEnum struct {
	AUTH_FAILURE                  DeleteGatewayResponseTypeV2RequestResponseType
	AUTH_HEADER_MISSING           DeleteGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_FAILURE            DeleteGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_CONF_FAILURE       DeleteGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_IDENTITIES_FAILURE DeleteGatewayResponseTypeV2RequestResponseType
	BACKEND_UNAVAILABLE           DeleteGatewayResponseTypeV2RequestResponseType
	BACKEND_TIMEOUT               DeleteGatewayResponseTypeV2RequestResponseType
	THROTTLED                     DeleteGatewayResponseTypeV2RequestResponseType
	UNAUTHORIZED                  DeleteGatewayResponseTypeV2RequestResponseType
	ACCESS_DENIED                 DeleteGatewayResponseTypeV2RequestResponseType
	NOT_FOUND                     DeleteGatewayResponseTypeV2RequestResponseType
	REQUEST_PARAMETERS_FAILURE    DeleteGatewayResponseTypeV2RequestResponseType
	DEFAULT_4_XX                  DeleteGatewayResponseTypeV2RequestResponseType
	DEFAULT_5_XX                  DeleteGatewayResponseTypeV2RequestResponseType
}

func GetDeleteGatewayResponseTypeV2RequestResponseTypeEnum() DeleteGatewayResponseTypeV2RequestResponseTypeEnum {
	return DeleteGatewayResponseTypeV2RequestResponseTypeEnum{
		AUTH_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "AUTH_FAILURE",
		},
		AUTH_HEADER_MISSING: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "AUTH_HEADER_MISSING",
		},
		AUTHORIZER_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_FAILURE",
		},
		AUTHORIZER_CONF_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_CONF_FAILURE",
		},
		AUTHORIZER_IDENTITIES_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_IDENTITIES_FAILURE",
		},
		BACKEND_UNAVAILABLE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "BACKEND_UNAVAILABLE",
		},
		BACKEND_TIMEOUT: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "BACKEND_TIMEOUT",
		},
		THROTTLED: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "THROTTLED",
		},
		UNAUTHORIZED: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "UNAUTHORIZED",
		},
		ACCESS_DENIED: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "ACCESS_DENIED",
		},
		NOT_FOUND: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "NOT_FOUND",
		},
		REQUEST_PARAMETERS_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "REQUEST_PARAMETERS_FAILURE",
		},
		DEFAULT_4_XX: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "DEFAULT_4XX",
		},
		DEFAULT_5_XX: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "DEFAULT_5XX",
		},
	}
}

func (c DeleteGatewayResponseTypeV2RequestResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteGatewayResponseTypeV2RequestResponseType) UnmarshalJSON(b []byte) error {
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
