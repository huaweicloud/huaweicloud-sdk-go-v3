package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRealTimeClipResponse Response Object
type CreateRealTimeClipResponse struct {

	// 截取的任务id，回调时会返回该任务id
	TaskId *string `json:"task_id,omitempty"`

	// 直播推流域名
	PublishDomain *string `json:"publish_domain,omitempty"`

	// 应用名
	App *string `json:"app,omitempty"`

	// 录制的流名
	Stream *string `json:"stream,omitempty"`

	// 录制完成文件对应的启动录制时间，UNIX时间戳
	StartTime *int32 `json:"start_time,omitempty"`

	// 录制完成文件对应的完成录制时间，UNIX时间戳
	EndTime *int32 `json:"end_time,omitempty"`

	// 粉丝截取响应的obs地址
	FileUrl *string `json:"file_url,omitempty"`

	Output *ObsInfo `json:"output,omitempty"`

	MetaData       *ObjectMetaData `json:"meta_data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateRealTimeClipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRealTimeClipResponse struct{}"
	}

	return strings.Join([]string{"CreateRealTimeClipResponse", string(data)}, " ")
}
