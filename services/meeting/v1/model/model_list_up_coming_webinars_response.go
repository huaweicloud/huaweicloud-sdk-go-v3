package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListUpComingWebinarsResponse struct {
	// 偏移量。

	Offset int32 `json:"offset"`
	// 查询个数。

	Limit int32 `json:"limit"`
	// 总记录数

	Count int64 `json:"count"`

	Data           *[]OpenWebinarUpcomingInfo `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListUpComingWebinarsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpComingWebinarsResponse struct{}"
	}

	return strings.Join([]string{"ListUpComingWebinarsResponse", string(data)}, " ")
}
