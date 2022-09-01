package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointParam struct {

	// 是否自动触发,true:自动触发，false：非自动触发。
	AutoTrigger *bool `json:"auto_trigger,omitempty" xml:"auto_trigger"`

	// 备份描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 是否增量备份，true:增量备份，false：非增量备份。
	Incremental *bool `json:"incremental,omitempty" xml:"incremental"`

	// 备份名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 待备份的资源id列表:uuid
	Resources *[]string `json:"resources,omitempty" xml:"resources"`

	// 资源详情
	ResourceDetails *[]Resource `json:"resource_details,omitempty" xml:"resource_details"`

	// 自动备份时的策略id
	PolicyId *string `json:"policy_id,omitempty" xml:"policy_id"`
}

func (o CheckpointParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointParam struct{}"
	}

	return strings.Join([]string{"CheckpointParam", string(data)}, " ")
}
