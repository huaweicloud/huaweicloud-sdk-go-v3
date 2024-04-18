package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetailsBody 详细信息结构体
type DetailsBody struct {

	// 变更前的容量，仅在变更规格时有值。
	OldCapacity *string `json:"old_capacity,omitempty"`

	// 变更后的容量，仅在变更规格时有值。
	NewCapacity *string `json:"new_capacity,omitempty"`

	// 是否开启公网访问，仅在开启公网访问时有值。
	EnablePublicIp *bool `json:"enable_public_ip,omitempty"`

	// 公网IP的ID，仅在开启公网访问时有值。
	PublicIpId *string `json:"public_ip_id,omitempty"`

	// 公网IP地址，仅在开启公网访问时有值。
	PublicIpAddress *string `json:"public_ip_address,omitempty"`

	// 是否开启SSL，仅在开启SSL时有值。
	EnableSsl *bool `json:"enable_ssl,omitempty"`

	// 变更前的缓存类型，仅在变更规格时有值。
	OldCacheMode *string `json:"old_cache_mode,omitempty"`

	// 变更后的缓存类型，仅在变更规格时有值。
	NewCacheMode *string `json:"new_cache_mode,omitempty"`

	// 变更前的规格参数，仅在变更规格时有值。
	OldResourceSpecCode *string `json:"old_resource_spec_code,omitempty"`

	// 变更后的规格参数，仅在变更规格时有值。
	NewResourceSpecCode *string `json:"new_resource_spec_code,omitempty"`

	// 变更前的副本数量，仅在变更规格时有值。
	OldReplicaNum *int32 `json:"old_replica_num,omitempty"`

	// 变更后的副本数量，仅在变更规格时有值。
	NewReplicaNum *int32 `json:"new_replica_num,omitempty"`

	// 变更前的缓存类型，仅在变更规格时有值。
	OldCacheType *string `json:"old_cache_type,omitempty"`

	// 变更后的规格类型，仅在变更规格时有值。
	NewCacheType *string `json:"new_cache_type,omitempty"`

	// 副本IP。
	ReplicaIp *string `json:"replica_ip,omitempty"`

	// 副本所在可用区。
	ReplicaAz *string `json:"replica_az,omitempty"`

	// 组名。
	GroupName *string `json:"group_name,omitempty"`

	// 旧端口。
	OldPort *int32 `json:"old_port,omitempty"`

	// 新端口。
	NewPort *int32 `json:"new_port,omitempty"`

	// 是否只是调整计费模式。
	IsOnlyAdjustCharging *bool `json:"is_only_adjust_charging,omitempty"`

	// 账号名称。
	AccountName *string `json:"account_name,omitempty"`

	// 源IP。
	SourceIp *string `json:"source_ip,omitempty"`

	// 目标IP。
	TargetIp *string `json:"target_ip,omitempty"`

	// 节点信息。
	NodeName *string `json:"node_name,omitempty"`

	// 重命名的指令。
	RenameCommands *[]string `json:"rename_commands,omitempty"`

	// 更新配置项的长度。
	UpdatedConfigLength *int32 `json:"updated_config_length,omitempty"`
}

func (o DetailsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetailsBody struct{}"
	}

	return strings.Join([]string{"DetailsBody", string(data)}, " ")
}
