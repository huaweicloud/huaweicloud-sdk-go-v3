package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SaveBackupDownloadPolicyRequestBody struct {

	// 备份下载开关。open表示打开备份下载开关，允许外网下载。close表示关闭备份下载开关，不允许外网下载。
	Action string `json:"action"`
}

func (o SaveBackupDownloadPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveBackupDownloadPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SaveBackupDownloadPolicyRequestBody", string(data)}, " ")
}
