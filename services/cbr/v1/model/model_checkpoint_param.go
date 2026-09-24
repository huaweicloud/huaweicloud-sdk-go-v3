package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointParam struct {

	// 是否自动触发,true：自动触发，false：非自动触发。
	AutoTrigger *bool `json:"auto_trigger,omitempty"`

	// 备份描述
	Description *string `json:"description,omitempty"`

	// 是否增量备份，true：增量备份，false：非增量备份。
	Incremental *bool `json:"incremental,omitempty"`

	// 备份名称
	Name *string `json:"name,omitempty"`

	// 待备份的资源id列表:uuid
	Resources *[]string `json:"resources,omitempty"`

	// 资源详情
	ResourceDetails *[]Resource `json:"resource_details,omitempty"`

	// 自动备份时的策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**： 手动备份的保留时长，单位为天。设置该参数后，备份副本将在保留时长到期后自动删除。用于为手动备份设置自动过期时间，避免手动备份堆积导致存储容量浪费。不设置此参数时，备份将永久保留。 **约束限制**： 当auto_trigger为true时不支持传此参数，自动备份的保留时间由关联的备份策略指定。auto_trigger不传或为false时支持指定此参数。 **取值范围**： -  1~36500：指定保留天数，备份将在创建时间 + 该天数后到期并自动删除。 - -1：永久保留，备份不会自动过期。  **默认取值**： -1 > 该特性目前处于公测阶段，部分Region可能无法使用
	RetentionDurationDays *int32 `json:"retention_duration_days,omitempty"`
}

func (o CheckpointParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointParam struct{}"
	}

	return strings.Join([]string{"CheckpointParam", string(data)}, " ")
}
