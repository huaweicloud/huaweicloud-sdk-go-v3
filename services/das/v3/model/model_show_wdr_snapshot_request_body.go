package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWdrSnapshotRequestBody 获取WDR快照列表请求体
type ShowWdrSnapshotRequestBody struct {

	// 开始时间（Unix timestamp），单位：毫秒
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ShowWdrSnapshotRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWdrSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"ShowWdrSnapshotRequestBody", string(data)}, " ")
}
