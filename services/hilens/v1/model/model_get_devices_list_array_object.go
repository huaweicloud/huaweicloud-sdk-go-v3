package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDevicesListArrayObject 设备实体列表
type GetDevicesListArrayObject struct {

	// 设备ID
	Id *string `json:"id,omitempty"`

	// 设备名称
	Name *string `json:"name,omitempty"`

	// 设备类型
	Type *string `json:"type,omitempty"`

	// 设备状态(0:离线;1:在线)
	Status *int32 `json:"status,omitempty"`

	// cpu核数
	Cpu *int32 `json:"cpu,omitempty"`

	// 内存大小
	Memory *int32 `json:"memory,omitempty"`

	// 操作系统
	Os *string `json:"os,omitempty"`

	// 固件名称
	FirmwareName *string `json:"firmware_name,omitempty"`

	// 固件版本
	FirmwareVersion *string `json:"firmware_version,omitempty"`

	// 固件状态(1:更新中，2：更新失败，3：更新成功)
	FirmwareStatus *int32 `json:"firmware_status,omitempty"`

	// 固件更新失败原因
	FirmwareCause *string `json:"firmware_cause,omitempty"`

	// 设备数据存储路径，该桶需要和当前region匹配
	Path *string `json:"path,omitempty"`

	// 设备数据存储路径更新状态(0:更新成功，1：更新中)
	PathUpdateStatus *int32 `json:"path_update_status,omitempty"`

	// 设备数据存储路径更新失败原因
	PathUpdateCause *string `json:"path_update_cause,omitempty"`

	// 创建时间（时间戳）
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间（时间戳）
	UpdateTime *int64 `json:"update_time,omitempty"`

	// IAM用户名
	UserTime *string `json:"user_time,omitempty"`

	// 计费资源码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 云服务计费类型
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 激活内容
	ActiveContent *string `json:"active_content,omitempty"`

	// 激活状态(0:未激活，1：已激活且付费，2：已激活且免费，3：付费到期，4：已激活使用SN码，5：已激活30天免费，6：免费到期)
	ActiveFlag *int32 `json:"active_flag,omitempty"`

	// 关联设备的主题消息推送的URN地址
	TopicUrn *string `json:"topic_urn,omitempty"`
}

func (o GetDevicesListArrayObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDevicesListArrayObject struct{}"
	}

	return strings.Join([]string{"GetDevicesListArrayObject", string(data)}, " ")
}
