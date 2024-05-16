package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostResultsRequest Request Object
type ListHostResultsRequest struct {

	// 主机ID
	HostId string `json:"host_id"`

	// 扫描ID
	ScanId string `json:"scan_id"`

	// 分页查询，偏移量，表示从此偏移量开始查询
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询，每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHostResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostResultsRequest struct{}"
	}

	return strings.Join([]string{"ListHostResultsRequest", string(data)}, " ")
}
