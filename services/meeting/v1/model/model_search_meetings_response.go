package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchMeetingsResponse Response Object
type SearchMeetingsResponse struct {

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

func (o SearchMeetingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchMeetingsResponse struct{}"
	}

	return strings.Join([]string{"SearchMeetingsResponse", string(data)}, " ")
}
