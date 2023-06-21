package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPortalInfosRequest struct {

	// 服务号名称。
	PubName *string `json:"pub_name,omitempty"`

	// 起始上线时间。格式为：yyyy-MM-ddTHH:mm:ssZ。
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 截止上线时间。格式为：yyyy-MM-ddTHH:mm:ssZ。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 主页状态。 - 1：未生效  - 2：已生效  - 3：已失效  - 4：服务号已冻结
	State *int32 `json:"state,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListPortalInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortalInfosRequest struct{}"
	}

	return strings.Join([]string{"ListPortalInfosRequest", string(data)}, " ")
}
