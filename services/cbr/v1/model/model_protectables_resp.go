package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProtectablesResp struct {

	// 子资源
	Children []interface{} `json:"children"`

	// 资源详情
	Detail *interface{} `json:"detail,omitempty"`

	// id
	Id string `json:"id"`

	// 名称
	Name string `json:"name"`

	Protectable *ProtectableResult `json:"protectable"`

	// 大小，单位GB
	Size *int32 `json:"size,omitempty"`

	// 资源状态
	Status *ProtectablesRespStatus `json:"status,omitempty"`

	// 待备份资源的类型, 云服务器: OS::Nova::Server, 云硬盘: OS::Cinder::Volume, 裸金属服务器: OS::Ironic::BareMetalServer, 线下本地服务器: OS::Native::Server, 弹性文件系统: OS::Sfs::Turbo, 云桌面：OS::Workspace::DesktopV2
	Type string `json:"type"`
}

func (o ProtectablesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectablesResp struct{}"
	}

	return strings.Join([]string{"ProtectablesResp", string(data)}, " ")
}

type ProtectablesRespStatus struct {
	value string
}

type ProtectablesRespStatusEnum struct {
	ACTIVE  ProtectablesRespStatus
	DELETED ProtectablesRespStatus
	ERROR   ProtectablesRespStatus
}

func GetProtectablesRespStatusEnum() ProtectablesRespStatusEnum {
	return ProtectablesRespStatusEnum{
		ACTIVE: ProtectablesRespStatus{
			value: "active",
		},
		DELETED: ProtectablesRespStatus{
			value: "deleted",
		},
		ERROR: ProtectablesRespStatus{
			value: "error",
		},
	}
}

func (c ProtectablesRespStatus) Value() string {
	return c.value
}

func (c ProtectablesRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProtectablesRespStatus) UnmarshalJSON(b []byte) error {
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
