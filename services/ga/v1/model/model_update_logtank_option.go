package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogtankOption 更新云日志的详细信息。
type UpdateLogtankOption struct {

	// 云日志服务提供的日志组ID。
	LogGroupId string `json:"log_group_id"`

	// 云日志服务提供的日志流ID。
	LogStreamId string `json:"log_stream_id"`
}

func (o UpdateLogtankOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankOption struct{}"
	}

	return strings.Join([]string{"UpdateLogtankOption", string(data)}, " ")
}
