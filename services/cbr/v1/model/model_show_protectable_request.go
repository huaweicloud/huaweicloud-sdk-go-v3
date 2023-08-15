package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowProtectableRequest Request Object
type ShowProtectableRequest struct {

	// 资源ID
	InstanceId string `json:"instance_id"`

	// 可保护性类型
	ProtectableType ShowProtectableRequestProtectableType `json:"protectable_type"`
}

func (o ShowProtectableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectableRequest struct{}"
	}

	return strings.Join([]string{"ShowProtectableRequest", string(data)}, " ")
}

type ShowProtectableRequestProtectableType struct {
	value string
}

type ShowProtectableRequestProtectableTypeEnum struct {
	SERVER ShowProtectableRequestProtectableType
	DISK   ShowProtectableRequestProtectableType
}

func GetShowProtectableRequestProtectableTypeEnum() ShowProtectableRequestProtectableTypeEnum {
	return ShowProtectableRequestProtectableTypeEnum{
		SERVER: ShowProtectableRequestProtectableType{
			value: "server",
		},
		DISK: ShowProtectableRequestProtectableType{
			value: "disk",
		},
	}
}

func (c ShowProtectableRequestProtectableType) Value() string {
	return c.value
}

func (c ShowProtectableRequestProtectableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProtectableRequestProtectableType) UnmarshalJSON(b []byte) error {
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
