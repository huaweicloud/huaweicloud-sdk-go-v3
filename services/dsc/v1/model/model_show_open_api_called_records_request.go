package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowOpenApiCalledRecordsRequest struct {

	// 分页大小，默认1000，最大2000。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 需要查询调用记录的URL，例如：/v1/{project_id}/sdg/database/watermark/embed。
	CalledUrl *string `json:"called_url,omitempty" xml:"called_url"`

	// 开始时间（Unix timestamp），单位：毫秒，例如：0
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time"`

	// 结束时间（Unix timestamp），单位：毫秒，例如：1638515803572
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time"`

	// 指定一个标识符。获取第一页时不用赋值，获取下一页时取上页查询结果的返回值。
	Marker *string `json:"marker,omitempty" xml:"marker"`
}

func (o ShowOpenApiCalledRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpenApiCalledRecordsRequest struct{}"
	}

	return strings.Join([]string{"ShowOpenApiCalledRecordsRequest", string(data)}, " ")
}
