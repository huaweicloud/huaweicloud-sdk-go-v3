package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHistoryWebinarsResponse struct {

	// 偏移量。
	Offset int32 `json:"offset" xml:"offset"`

	// 查询个数。
	Limit int32 `json:"limit" xml:"limit"`

	// 总记录数
	Count int64 `json:"count" xml:"count"`

	Data           *[]OpenWebinarHistoryInfo `json:"data,omitempty" xml:"data"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListHistoryWebinarsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryWebinarsResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryWebinarsResponse", string(data)}, " ")
}
