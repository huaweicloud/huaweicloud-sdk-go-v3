package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchRecordingsResponse struct {

	// 第几条。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页的记录数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 总记录数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 录播文件列表。
	Data           *[]RecordResultDo `json:"data,omitempty" xml:"data"`
	HttpStatusCode int               `json:"-"`
}

func (o SearchRecordingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchRecordingsResponse struct{}"
	}

	return strings.Join([]string{"SearchRecordingsResponse", string(data)}, " ")
}
