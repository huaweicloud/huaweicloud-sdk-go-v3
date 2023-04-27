package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowNextflowTaskLogResponse struct {

	// 任务日志条数
	Count *int32 `json:"count,omitempty"`

	// 日志内容列表
	Logs *[]string `json:"logs,omitempty"`

	// 日志下载链接
	DownloadUrl    *string `json:"download_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowNextflowTaskLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNextflowTaskLogResponse struct{}"
	}

	return strings.Join([]string{"ShowNextflowTaskLogResponse", string(data)}, " ")
}
