package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditInstancesLastResponse Response Object
type ListAuditInstancesLastResponse struct {

	// 实例信息列表
	Instances *[]AuditInstanceBean `json:"instances,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditInstancesLastResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditInstancesLastResponse struct{}"
	}

	return strings.Join([]string{"ListAuditInstancesLastResponse", string(data)}, " ")
}
