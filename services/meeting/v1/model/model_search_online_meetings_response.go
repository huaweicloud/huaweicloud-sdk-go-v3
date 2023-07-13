package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchOnlineMeetingsResponse Response Object
type SearchOnlineMeetingsResponse struct {

	// 查询偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 每页的记录数。
	Limit *int32 `json:"limit,omitempty"`

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 会议列表。
	Data           *[]ConferenceInfo `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o SearchOnlineMeetingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchOnlineMeetingsResponse struct{}"
	}

	return strings.Join([]string{"SearchOnlineMeetingsResponse", string(data)}, " ")
}
