package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmsDetailsRequest Request Object
type ListSmsDetailsRequest struct {

	// 分页查询时每页显示的记录数，默认值为10，取值范围为10-500的整数
	Limit *int64 `json:"limit,omitempty"`

	// 分页查询时的页码数，默认值为1，取值范围为1-1000000的整数
	Offset *int64 `json:"offset,omitempty"`

	// 容器ID
	Cid *string `json:"cid,omitempty"`

	// 开始时间
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 结束时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`
}

func (o ListSmsDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmsDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSmsDetailsRequest", string(data)}, " ")
}
