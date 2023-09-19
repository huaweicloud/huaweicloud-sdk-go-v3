package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogMappingStreamInfo struct {

	// 源日志流ID
	SourceLogStreamId *string `json:"source_log_stream_id,omitempty"`

	// 目标日志流ID
	TargetLogStreamId *string `json:"target_log_stream_id,omitempty"`

	// 目标日志流名称
	TargetLogStreamName *string `json:"target_log_stream_name,omitempty"`

	// 目标日志流EPS ID
	TargetLogStreamEpsId *string `json:"target_log_stream_eps_id,omitempty"`

	// 目标日志流ttl
	TargetLogStreamTtl *int32 `json:"target_log_stream_ttl,omitempty"`
}

func (o LogMappingStreamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogMappingStreamInfo struct{}"
	}

	return strings.Join([]string{"LogMappingStreamInfo", string(data)}, " ")
}
