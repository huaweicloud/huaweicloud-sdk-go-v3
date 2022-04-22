package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListInvoicesRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 发票申请开始时间。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	StartTime string `json:"start_time"`

	// 发票申请结束时间。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	EndTime string `json:"end_time"`

	// 页码。
	Offset int32 `json:"offset"`

	// 每页大小。
	Limit int32 `json:"limit"`
}

func (o ListInvoicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInvoicesRequest struct{}"
	}

	return strings.Join([]string{"ListInvoicesRequest", string(data)}, " ")
}
