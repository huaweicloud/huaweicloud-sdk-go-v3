package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessClientRequestBody 创建接入客户端的请求信息
type AccessClientRequestBody struct {

	// 客户端名称
	Name string `json:"name"`

	// 接入模式， SYSTEM：系统默认模式，由系统自动创建vpcep连接，也是推荐方式。该模式下vpc_id、subnet_id必填。 CUSTOM：定制模式，由外部服务自行创建vpcep连接，适用于跨租户场景等。该模式下access_connection必填。
	AccessMode *AccessClientRequestBodyAccessMode `json:"access_mode,omitempty"`

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 接入连接列表
	AccessConnections *[]AccessConnectionRequestBody `json:"access_connections,omitempty"`
}

func (o AccessClientRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessClientRequestBody struct{}"
	}

	return strings.Join([]string{"AccessClientRequestBody", string(data)}, " ")
}

type AccessClientRequestBodyAccessMode struct {
	value string
}

type AccessClientRequestBodyAccessModeEnum struct {
	SYSTEM AccessClientRequestBodyAccessMode
	CUSTOM AccessClientRequestBodyAccessMode
}

func GetAccessClientRequestBodyAccessModeEnum() AccessClientRequestBodyAccessModeEnum {
	return AccessClientRequestBodyAccessModeEnum{
		SYSTEM: AccessClientRequestBodyAccessMode{
			value: "SYSTEM",
		},
		CUSTOM: AccessClientRequestBodyAccessMode{
			value: "CUSTOM",
		},
	}
}

func (c AccessClientRequestBodyAccessMode) Value() string {
	return c.value
}

func (c AccessClientRequestBodyAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessClientRequestBodyAccessMode) UnmarshalJSON(b []byte) error {
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
