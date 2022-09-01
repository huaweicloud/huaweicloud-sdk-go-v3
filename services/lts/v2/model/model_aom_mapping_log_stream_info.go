package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AomMappingLogStreamInfo struct {

	// 日志组id
	TargetLogGroupId string `json:"target_log_group_id" xml:"target_log_group_id"`

	// 目标日志组名称。
	TargetLogGroupName string `json:"target_log_group_name" xml:"target_log_group_name"`

	// 日志流id
	TargetLogStreamId string `json:"target_log_stream_id" xml:"target_log_stream_id"`

	// 目标日志组名称。
	TargetLogStreamName string `json:"target_log_stream_name" xml:"target_log_stream_name"`
}

func (o AomMappingLogStreamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AomMappingLogStreamInfo struct{}"
	}

	return strings.Join([]string{"AomMappingLogStreamInfo", string(data)}, " ")
}
