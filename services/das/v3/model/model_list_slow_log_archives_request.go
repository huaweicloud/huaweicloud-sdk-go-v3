package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowLogArchivesRequest Request Object
type ListSlowLogArchivesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 当前页码
	CurPage *int32 `json:"cur_page,omitempty"`

	// 页大小
	PerPage *int32 `json:"per_page,omitempty"`

	// 开始时间（Unix时间戳，毫秒）
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix时间戳，毫秒）
	EndTime int64 `json:"end_time"`
}

func (o ListSlowLogArchivesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogArchivesRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogArchivesRequest", string(data)}, " ")
}
