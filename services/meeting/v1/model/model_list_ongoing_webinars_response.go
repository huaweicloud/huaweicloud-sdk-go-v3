package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOngoingWebinarsResponse Response Object
type ListOngoingWebinarsResponse struct {

	// 偏移量。
	Offset int32 `json:"offset"`

	// 每页的记录数。
	Limit int32 `json:"limit"`

	// 总记录数。
	Count int64 `json:"count"`

	// 正在召开网络研讨会信息列表。
	Data           *[]OpenWebinarOngoingInfo `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListOngoingWebinarsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOngoingWebinarsResponse struct{}"
	}

	return strings.Join([]string{"ListOngoingWebinarsResponse", string(data)}, " ")
}
