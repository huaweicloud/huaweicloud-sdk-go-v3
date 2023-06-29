package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistorySqlRequest Request Object
type ListHistorySqlRequest struct {

	// 日志组id
	LogGroupId string `json:"log_group_id"`

	// 日志流id
	LogStreamId string `json:"log_stream_id"`
}

func (o ListHistorySqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistorySqlRequest struct{}"
	}

	return strings.Join([]string{"ListHistorySqlRequest", string(data)}, " ")
}
