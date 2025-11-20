package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAssetFileResponse Response Object
type DownloadAssetFileResponse struct {

	// **参数解释**： 导出任务ID **取值范围**： 字符长度0-128位
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadAssetFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAssetFileResponse struct{}"
	}

	return strings.Join([]string{"DownloadAssetFileResponse", string(data)}, " ")
}
