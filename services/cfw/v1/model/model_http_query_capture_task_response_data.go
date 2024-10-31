package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpQueryCaptureTaskResponseData struct {

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 抓包任务总数
	Total *int64 `json:"total,omitempty"`

	// 抓包任务列表
	Records *[]CaptureTaskVo `json:"records,omitempty"`
}

func (o HttpQueryCaptureTaskResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpQueryCaptureTaskResponseData struct{}"
	}

	return strings.Join([]string{"HttpQueryCaptureTaskResponseData", string(data)}, " ")
}
