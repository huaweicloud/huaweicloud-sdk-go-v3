package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditInstancesRequest Request Object
type ListAuditInstancesRequest struct {

	// 偏移量，默认0。
	Offset *string `json:"offset,omitempty"`

	// 查询记录数，默认100，最大1000。
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAuditInstancesRequest", string(data)}, " ")
}
