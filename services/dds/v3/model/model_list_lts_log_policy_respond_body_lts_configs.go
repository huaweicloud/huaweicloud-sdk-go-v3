package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListLtsLogPolicyRespondBodyLtsConfigs struct {
	LogType *LtsLogType `json:"log_type,omitempty"`

	// 云日志服务LTS日志组ID。
	LtsGroupId *string `json:"lts_group_id,omitempty"`

	// 云日志服务LTS日志流ID。
	LtsStreamId *string `json:"lts_stream_id,omitempty"`

	// 是否上传。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o ListLtsLogPolicyRespondBodyLtsConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsLogPolicyRespondBodyLtsConfigs struct{}"
	}

	return strings.Join([]string{"ListLtsLogPolicyRespondBodyLtsConfigs", string(data)}, " ")
}
