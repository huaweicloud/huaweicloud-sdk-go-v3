package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpComingWebinarsResponse Response Object
type ListUpComingWebinarsResponse struct {

	// 偏移量。
	Offset int32 `json:"offset"`

	// 每页的记录数。
	Limit int32 `json:"limit"`

	// 总记录数。
	Count int64 `json:"count"`

	// 即将召开研讨会信息列表。
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
