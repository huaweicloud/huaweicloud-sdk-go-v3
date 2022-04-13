package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchAttendanceRecordsOfHisMeetingResponse struct {
	// 第几条。

	Offset *int32 `json:"offset,omitempty"`
	// 每页的记录数。

	Limit *int32 `json:"limit,omitempty"`
	// 总记录数。

	Count *int32 `json:"count,omitempty"`
	// 与会者列表。

	Data           *[]ConfAttendeeRecordInfo `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o SearchAttendanceRecordsOfHisMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAttendanceRecordsOfHisMeetingResponse struct{}"
	}

	return strings.Join([]string{"SearchAttendanceRecordsOfHisMeetingResponse", string(data)}, " ")
}
