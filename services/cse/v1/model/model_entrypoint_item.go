package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EntrypointItem struct {

	// 微服务引擎专享版组件的ipv4主接入地址
	MasterEntrypoint *string `json:"master_entrypoint,omitempty" xml:"master_entrypoint"`

	// 微服务引擎专享版组件的ipv6主接入地址
	MasterEntrypointIpv6 *string `json:"master_entrypoint_ipv6,omitempty" xml:"master_entrypoint_ipv6"`

	// 微服务引擎专享版组件的ipv4备接入地址
	SlaveEntrypoint *string `json:"slave_entrypoint,omitempty" xml:"slave_entrypoint"`

	// 微服务引擎专享版组件的ipv6备接入地址
	SlaveEntrypointIpv6 *string `json:"slave_entrypoint_ipv6,omitempty" xml:"slave_entrypoint_ipv6"`

	// 微服务引擎专享版组件类型
	Type *EntrypointItemType `json:"type,omitempty" xml:"type"`
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
