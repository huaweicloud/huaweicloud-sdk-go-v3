package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchMeetingFileListResponse struct {
	// 页面起始页，从0开始

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的条目数量。 默认值：10。

	Limit *int32 `json:"limit,omitempty"`
	// 总数量。

	Count *int32 `json:"count,omitempty"`
	// IdeaHub、终端“查看会议纪要列表”响应

	Data           *[]ListMeetingFileResponseDto `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o SearchMeetingFileListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchMeetingFileListResponse struct{}"
	}

	return strings.Join([]string{"SearchMeetingFileListResponse", string(data)}, " ")
}
