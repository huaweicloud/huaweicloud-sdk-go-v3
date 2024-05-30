package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncStatusStatisticVo struct {

	// 成功数，填写String类型替代Long类型。
	Success *string `json:"success,omitempty"`

	// 失败数，填写String类型替代Long类型。
	Failed *string `json:"failed,omitempty"`

	// 同步中，填写String类型替代Long类型。
	Running *string `json:"running,omitempty"`

	// 未同步数，填写String类型替代Long类型。
	Other *string `json:"other,omitempty"`
}

func (o SyncStatusStatisticVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncStatusStatisticVo struct{}"
	}

	return strings.Join([]string{"SyncStatusStatisticVo", string(data)}, " ")
}
