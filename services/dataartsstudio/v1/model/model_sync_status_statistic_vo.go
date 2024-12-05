package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncStatusStatisticVo struct {

	// 成功数，ID字符串。
	Success *string `json:"success,omitempty"`

	// 失败数，ID字符串。
	Failed *string `json:"failed,omitempty"`

	// 同步中，ID字符串。
	Running *string `json:"running,omitempty"`

	// 未同步数，ID字符串。
	Other *string `json:"other,omitempty"`
}

func (o SyncStatusStatisticVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncStatusStatisticVo struct{}"
	}

	return strings.Join([]string{"SyncStatusStatisticVo", string(data)}, " ")
}
