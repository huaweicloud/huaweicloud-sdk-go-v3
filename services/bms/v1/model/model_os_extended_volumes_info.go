package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// os-extended-volumes:volumes_attached数据结构说明
type OsExtendedVolumesInfo struct {
	// 磁盘ID，格式为UUID

	Id *string `json:"id,omitempty"`
	// 删裸金属服务器时是否一并删除该磁盘。true：是false：否

	DeleteOnTermination *string `json:"delete_on_termination,omitempty"`
	// 启动标识，“0”代表启动盘，“-1“代表非启动盘

	BootIndex *OsExtendedVolumesInfoBootIndex `json:"bootIndex,omitempty"`
	// 磁盘设备名称，例如“/dev/sdb”。

	Device *string `json:"device,omitempty"`
}

func (o OsExtendedVolumesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsExtendedVolumesInfo struct{}"
	}

	return strings.Join([]string{"OsExtendedVolumesInfo", string(data)}, " ")
}

type OsExtendedVolumesInfoBootIndex struct {
	value string
}

type OsExtendedVolumesInfoBootIndexEnum struct {
	E_0 OsExtendedVolumesInfoBootIndex
	E_1 OsExtendedVolumesInfoBootIndex
}

func GetOsExtendedVolumesInfoBootIndexEnum() OsExtendedVolumesInfoBootIndexEnum {
	return OsExtendedVolumesInfoBootIndexEnum{
		E_0: OsExtendedVolumesInfoBootIndex{
			value: "0",
		},
		E_1: OsExtendedVolumesInfoBootIndex{
			value: "-1",
		},
	}
}

func (c OsExtendedVolumesInfoBootIndex) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OsExtendedVolumesInfoBootIndex) UnmarshalJSON(b []byte) error {
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
