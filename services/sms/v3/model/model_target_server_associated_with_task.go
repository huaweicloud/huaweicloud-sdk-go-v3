package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TargetServerAssociatedWithTask 任务关联的目的端信息
type TargetServerAssociatedWithTask struct {

	// 目的端在SMS数据库中的ID
	Id *string `json:"id,omitempty"`

	// 目的端虚机ID
	VmId *string `json:"vm_id,omitempty"`

	// 目的端服务器名称
	Name *string `json:"name,omitempty"`

	// 目的端服务器IP
	Ip *string `json:"ip,omitempty"`

	// 目的端服务器的OS类型 WINDOWS:WINDOWS系统 LINUX:LINUX系统
	OsType *TargetServerAssociatedWithTaskOsType `json:"os_type,omitempty"`

	// 操作系统版本
	OsVersion *string `json:"os_version,omitempty"`
}

func (o TargetServerAssociatedWithTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetServerAssociatedWithTask struct{}"
	}

	return strings.Join([]string{"TargetServerAssociatedWithTask", string(data)}, " ")
}

type TargetServerAssociatedWithTaskOsType struct {
	value string
}

type TargetServerAssociatedWithTaskOsTypeEnum struct {
	WINDOWS TargetServerAssociatedWithTaskOsType
	LINUX   TargetServerAssociatedWithTaskOsType
}

func GetTargetServerAssociatedWithTaskOsTypeEnum() TargetServerAssociatedWithTaskOsTypeEnum {
	return TargetServerAssociatedWithTaskOsTypeEnum{
		WINDOWS: TargetServerAssociatedWithTaskOsType{
			value: "WINDOWS",
		},
		LINUX: TargetServerAssociatedWithTaskOsType{
			value: "LINUX",
		},
	}
}

func (c TargetServerAssociatedWithTaskOsType) Value() string {
	return c.value
}

func (c TargetServerAssociatedWithTaskOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetServerAssociatedWithTaskOsType) UnmarshalJSON(b []byte) error {
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
