package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ServersInfoType 数据结构说明
type ServersInfoType struct {

	// 重启类型：SOFT：普通重启（不生效）。HARD：强制重启（默认）。
	Type ServersInfoTypeType `json:"type"`

	// 裸金属服务器ID列表，详情请参见表3 servers字段数据结构说明。
	Servers []ServersList `json:"servers"`
}

func (o ServersInfoType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServersInfoType struct{}"
	}

	return strings.Join([]string{"ServersInfoType", string(data)}, " ")
}

type ServersInfoTypeType struct {
	value string
}

type ServersInfoTypeTypeEnum struct {
	SOFT ServersInfoTypeType
	HARD ServersInfoTypeType
}

func GetServersInfoTypeTypeEnum() ServersInfoTypeTypeEnum {
	return ServersInfoTypeTypeEnum{
		SOFT: ServersInfoTypeType{
			value: "SOFT",
		},
		HARD: ServersInfoTypeType{
			value: "HARD",
		},
	}
}

func (c ServersInfoTypeType) Value() string {
	return c.value
}

func (c ServersInfoTypeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServersInfoTypeType) UnmarshalJSON(b []byte) error {
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
