package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointParam struct {

	// 是否自动触发,true:自动触发，false：非自动触发。
	AutoTrigger *bool `json:"auto_trigger,omitempty"`

	// 备份描述
	Description *string `json:"description,omitempty"`

	// 是否增量备份，true:增量备份，false：非增量备份。
	Incremental *bool `json:"incremental,omitempty"`

	// 备份名称
	Name *string `json:"name,omitempty"`

	// 待备份的资源id列表:uuid
	Resources *[]string `json:"resources,omitempty"`

	// 资源详情
	ResourceDetails *[]Resource `json:"resource_details,omitempty"`

	// 自动备份时的策略id
	PolicyId *string `json:"policy_id,omitempty"`
}

func (o CheckpointParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointParam struct{}"
	}

	return strings.Join([]string{"CheckpointParam", string(data)}, " ")
}
