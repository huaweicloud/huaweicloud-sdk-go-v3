package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScreenRecordsTrafficLimitConfigRequest Request Object
type ListScreenRecordsTrafficLimitConfigRequest struct {

	// 站点ID。
	SiteId *string `json:"site_id,omitempty"`

	// 用于分页查询，返回录屏记录数量的限制。默认100。范围0~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListScreenRecordsTrafficLimitConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScreenRecordsTrafficLimitConfigRequest struct{}"
	}

	return strings.Join([]string{"ListScreenRecordsTrafficLimitConfigRequest", string(data)}, " ")
}
