package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBanUrlRequest Request Object
type ListBanUrlRequest struct {

	// 查询起始时间戳（毫秒），需与结束时间戳同时指定。
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询结束时间戳（毫秒），需与开始时间戳同时指定。
	EndTime *int64 `json:"end_time,omitempty"`

	// 封禁的url所显示单页最大数量，取值范围为1-50。page_size和page_number必须同时传值。默认值10.
	PageSize *int32 `json:"page_size,omitempty"`

	// 封禁的url当前查询为第几页，取值范围为1-65535。默认值1
	PageNumber *int32 `json:"page_number,omitempty"`

	// 封禁的url地址。
	Url *string `json:"url,omitempty"`
}

func (o ListBanUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBanUrlRequest struct{}"
	}

	return strings.Join([]string{"ListBanUrlRequest", string(data)}, " ")
}
