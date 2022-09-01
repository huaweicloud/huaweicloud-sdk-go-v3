package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLogStreamsRequest struct {

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty" xml:"log_group_name"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty" xml:"log_stream_name"`

	// 查询游标，初始传入0，后续从上一次的返回值中获取
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页数据量，最大值为100
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListLogStreamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamsRequest struct{}"
	}

	return strings.Join([]string{"ListLogStreamsRequest", string(data)}, " ")
}
