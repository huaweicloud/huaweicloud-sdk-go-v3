package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CheckpointResourceResp struct {

	// 资源附加信息
	ExtraInfo *string `json:"extra_info,omitempty"`

	// 待备份资源id
	Id string `json:"id"`

	// 待备份资源名称
	Name string `json:"name"`

	// 保护状态。available（可用），error（错误），protecting（备份中），restoring（恢复中），removing（删除中）。
	ProtectStatus *CheckpointResourceRespProtectStatus `json:"protect_status,omitempty"`

	// 资源已分配容量,单位为GB
	ResourceSize *string `json:"resource_size,omitempty"`

	// 待备份资源的类型: OS::Nova::Server, OS::Cinder::Volume, OS::Ironic::BareMetalServer, OS::Native::Server, OS::Sfs::Turbo, OS::Workspace::DesktopV2
	Type string `json:"type"`

	// 副本大小
	BackupSize *string `json:"backup_size,omitempty"`

	// 副本数量
	BackupCount *string `json:"backup_count,omitempty"`
}

func (o CheckpointResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointResourceResp struct{}"
	}

	return strings.Join([]string{"CheckpointResourceResp", string(data)}, " ")
}

type CheckpointResourceRespProtectStatus struct {
	value string
}

type CheckpointResourceRespProtectStatusEnum struct {
	AVAILABLE  CheckpointResourceRespProtectStatus
	ERROR      CheckpointResourceRespProtectStatus
	PROTECTING CheckpointResourceRespProtectStatus
	RESTORING  CheckpointResourceRespProtectStatus
	REMOVING   CheckpointResourceRespProtectStatus
}

func GetCheckpointResourceRespProtectStatusEnum() CheckpointResourceRespProtectStatusEnum {
	return CheckpointResourceRespProtectStatusEnum{
		AVAILABLE: CheckpointResourceRespProtectStatus{
			value: "available",
		},
		ERROR: CheckpointResourceRespProtectStatus{
			value: "error",
		},
		PROTECTING: CheckpointResourceRespProtectStatus{
			value: "protecting",
		},
		RESTORING: CheckpointResourceRespProtectStatus{
			value: "restoring",
		},
		REMOVING: CheckpointResourceRespProtectStatus{
			value: "removing",
		},
	}
}

func (c CheckpointResourceRespProtectStatus) Value() string {
	return c.value
}

func (c CheckpointResourceRespProtectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckpointResourceRespProtectStatus) UnmarshalJSON(b []byte) error {
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
