package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReplicationProfilesResult 发布订配置文件响应信息。
type ListReplicationProfilesResult struct {

	// 配置文件id。
	ProfileId *string `json:"profile_id,omitempty"`

	// 配置文件名。
	ProfileName *string `json:"profile_name,omitempty"`

	// 代理类型。  snapshot：快照代理 log_reader：日志读取器代理 distribution：分发代理 merge:合并代理 queue_reader：队列读取器代理。
	AgentType *string `json:"agent_type,omitempty"`

	// 配置文件说明。
	Description *string `json:"description,omitempty"`

	// 是否为默认配置。  true：是默认配置。 false：不是默认配置。
	IsDefProfile *bool `json:"is_def_profile,omitempty"`
}

func (o ListReplicationProfilesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplicationProfilesResult struct{}"
	}

	return strings.Join([]string{"ListReplicationProfilesResult", string(data)}, " ")
}
