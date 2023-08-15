package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditInstancesResponse Response Object
type ListAuditInstancesResponse struct {

	// 实例信息列表
	Servers *[]AuditInstanceListBean `json:"servers,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListAuditInstancesResponse", string(data)}, " ")
}
