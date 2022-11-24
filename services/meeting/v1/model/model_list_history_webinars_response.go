package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHistoryWebinarsResponse struct {

	// 偏移量。
	Offset int32 `json:"offset"`

	// 每页的记录数。
	Limit int32 `json:"limit"`

	// 总记录数。
	Count int64 `json:"count"`

	// 历史网络研讨会信息列表。
	Data           *[]OpenWebinarHistoryInfo `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListHistoryWebinarsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryWebinarsResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryWebinarsResponse", string(data)}, " ")
}
