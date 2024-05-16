package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostsRequest Request Object
type ListHostsRequest struct {

	// 主机组id
	GroupId *string `json:"group_id,omitempty"`

	// 分页查询，偏移量，表示从此偏移量开始查询
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询，每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsRequest struct{}"
	}

	return strings.Join([]string{"ListHostsRequest", string(data)}, " ")
}
