package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OsStopBodyType 数据结构说明
type OsStopBodyType struct {

	// 关机类型：SOFT：普通关机（不生效）。HARD：强制关机（默认）。
	Type *OsStopBodyTypeType `json:"type,omitempty"`

	// 裸金属服务器ID列表，详情请参见表3 servers字段数据结构说明。
	Servers []ServersList `json:"servers"`
}

func (o OsStopBodyType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsStopBodyType struct{}"
	}

	return strings.Join([]string{"OsStopBodyType", string(data)}, " ")
}

type OsStopBodyTypeType struct {
	value string
}

type OsStopBodyTypeTypeEnum struct {
	SOFT OsStopBodyTypeType
	HARD OsStopBodyTypeType
}

func GetOsStopBodyTypeTypeEnum() OsStopBodyTypeTypeEnum {
	return OsStopBodyTypeTypeEnum{
		SOFT: OsStopBodyTypeType{
			value: "SOFT",
		},
		HARD: OsStopBodyTypeType{
			value: "HARD",
		},
	}
}

func (c OsStopBodyTypeType) Value() string {
	return c.value
}

func (c OsStopBodyTypeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OsStopBodyTypeType) UnmarshalJSON(b []byte) error {
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
