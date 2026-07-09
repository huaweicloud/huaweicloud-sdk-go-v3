package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditInstancesNewResponse Response Object
type ListAuditInstancesNewResponse struct {

	// 实例列表
	Instances *[]ResponseAuditV2 `json:"instances,omitempty"`

	// 总记录数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditInstancesNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditInstancesNewResponse struct{}"
	}

	return strings.Join([]string{"ListAuditInstancesNewResponse", string(data)}, " ")
}
