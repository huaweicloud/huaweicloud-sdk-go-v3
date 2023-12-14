package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsLogsRequest Request Object
type ListLtsLogsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListLtsLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsLogsRequest struct{}"
	}

	return strings.Join([]string{"ListLtsLogsRequest", string(data)}, " ")
}
