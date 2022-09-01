package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchOnlineMeetingsResponse struct {

	// 第几条。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页的记录数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 总记录数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 会议信息列表。
	Data           *[]ConferenceInfo `json:"data,omitempty" xml:"data"`
	HttpStatusCode int               `json:"-"`
}

func (o SearchOnlineMeetingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchOnlineMeetingsResponse struct{}"
	}

	return strings.Join([]string{"SearchOnlineMeetingsResponse", string(data)}, " ")
}
