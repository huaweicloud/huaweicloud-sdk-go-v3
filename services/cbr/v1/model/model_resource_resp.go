package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response body of resource
type ResourceResp struct {
	ExtraInfo *ResourceExtraInfo `json:"extra_info,omitempty"`
	// 待备份资源id

	Id string `json:"id"`
	// 待备份资源名称

	Name string `json:"name"`
	// 保护状态

	ProtectStatus *ResourceRespProtectStatus `json:"protect_status,omitempty"`
	// 资源已分配容量,单位为GB

	Size *int32 `json:"size,omitempty"`
	// 待备份资源的类型, 云服务器: OS::Nova::Server, 云硬盘: OS::Cinder::Volume, 裸金属服务器: OS::Ironic::BareMetalServer, 线下本地服务器: OS::Native::Server, 弹性文件系统: OS::Sfs::Turbo

	Type string `json:"type"`
	// 副本大小

	BackupSize *int32 `json:"backup_size,omitempty"`
	// 副本数量

	BackupCount *int32 `json:"backup_count,omitempty"`
}

func (o ResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceResp struct{}"
	}

	return strings.Join([]string{"ResourceResp", string(data)}, " ")
}

type ResourceRespProtectStatus struct {
	value string
}

type ResourceRespProtectStatusEnum struct {
	AVAILABLE  ResourceRespProtectStatus
	ERROR      ResourceRespProtectStatus
	PROTECTING ResourceRespProtectStatus
	RESTORING  ResourceRespProtectStatus
	REMOVING   ResourceRespProtectStatus
}

func GetResourceRespProtectStatusEnum() ResourceRespProtectStatusEnum {
	return ResourceRespProtectStatusEnum{
		AVAILABLE: ResourceRespProtectStatus{
			value: "available",
		},
		ERROR: ResourceRespProtectStatus{
			value: "error",
		},
		PROTECTING: ResourceRespProtectStatus{
			value: "protecting",
		},
		RESTORING: ResourceRespProtectStatus{
			value: "restoring",
		},
		REMOVING: ResourceRespProtectStatus{
			value: "removing",
		},
	}
}

func (c ResourceRespProtectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceRespProtectStatus) UnmarshalJSON(b []byte) error {
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
