package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKillProcessHistoryRequest Request Object
type ListKillProcessHistoryRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 开始时间（Unix timestamp，毫秒）
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix timestamp，毫秒）
	EndTime int64 `json:"end_time"`

	// 页数
	PageNum int32 `json:"page_num"`

	// 页大小
	PageSize int32 `json:"page_size"`
}

func (o ListKillProcessHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKillProcessHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListKillProcessHistoryRequest", string(data)}, " ")
}
