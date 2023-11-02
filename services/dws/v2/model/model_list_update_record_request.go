package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpdateRecordRequest Request Object
type ListUpdateRecordRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 条目数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListUpdateRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpdateRecordRequest struct{}"
	}

	return strings.Join([]string{"ListUpdateRecordRequest", string(data)}, " ")
}
