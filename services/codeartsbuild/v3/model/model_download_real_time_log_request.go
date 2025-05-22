package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadRealTimeLogRequest Request Object
type DownloadRealTimeLogRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 构建任务的构建编号，从1开始，每次构建递增1
	BuildNo int32 `json:"build_no"`

	// 偏移量，传入前一次请求返回的offset
	Offset int32 `json:"offset"`

	// 可控制返回内容长度，默认值为1000000
	Length *int32 `json:"length,omitempty"`
}

func (o DownloadRealTimeLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadRealTimeLogRequest struct{}"
	}

	return strings.Join([]string{"DownloadRealTimeLogRequest", string(data)}, " ")
}
