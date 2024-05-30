package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EntrypointItem struct {

	// 微服务引擎组件的ipv4主接入地址
	MasterEntrypoint *string `json:"masterEntrypoint,omitempty"`

	// 微服务引擎组件的ipv6主接入地址
	MasterEntrypointIpv6 *string `json:"masterEntrypointIpv6,omitempty"`

	// 微服务引擎组件的ipv4备接入地址
	SlaveEntrypoint *string `json:"slaveEntrypoint,omitempty"`

	// 微服务引擎组件的ipv6备接入地址
	SlaveEntrypointIpv6 *string `json:"slaveEntrypointIpv6,omitempty"`

	// 微服务引擎组件类型
	Type *EntrypointItemType `json:"type,omitempty"`
}

func (o EntrypointItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntrypointItem struct{}"
	}

	return strings.Join([]string{"EntrypointItem", string(data)}, " ")
}

type EntrypointItemType struct {
	value string
}

type EntrypointItemTypeEnum struct {
	REGISTRY EntrypointItemType
	SERVICE  EntrypointItemType
}

func GetEntrypointItemTypeEnum() EntrypointItemTypeEnum {
	return EntrypointItemTypeEnum{
		REGISTRY: EntrypointItemType{
			value: "REGISTRY",
		},
		SERVICE: EntrypointItemType{
			value: "SERVICE",
		},
	}
}

func (c EntrypointItemType) Value() string {
	return c.value
}

func (c EntrypointItemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EntrypointItemType) UnmarshalJSON(b []byte) error {
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
