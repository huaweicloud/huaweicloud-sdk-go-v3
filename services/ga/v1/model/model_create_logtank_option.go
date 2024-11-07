package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogtankOption 创建云日志的详细信息。
type CreateLogtankOption struct {
	ResourceType *LogtankResourceType `json:"resource_type"`

	// 开启云日志的资源ID。
	ResourceId string `json:"resource_id"`

	// 云日志服务提供的日志组ID。
	LogGroupId string `json:"log_group_id"`

	// 云日志服务提供的日志流ID。
	LogStreamId string `json:"log_stream_id"`
}

func (o CreateLogtankOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankOption struct{}"
	}

	return strings.Join([]string{"CreateLogtankOption", string(data)}, " ")
}
