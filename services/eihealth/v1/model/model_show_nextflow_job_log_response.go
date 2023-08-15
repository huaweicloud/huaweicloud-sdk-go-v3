package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNextflowJobLogResponse Response Object
type ShowNextflowJobLogResponse struct {

	// 作业日志条数
	Count *int32 `json:"count,omitempty"`

	// 日志内容列表
	Logs *[]string `json:"logs,omitempty"`

	// 日志下载链接
	DownloadUrl    *string `json:"download_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowNextflowJobLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNextflowJobLogResponse struct{}"
	}

	return strings.Join([]string{"ShowNextflowJobLogResponse", string(data)}, " ")
}
