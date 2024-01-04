package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditionFile struct {

	// 试听文件下载链接，有效期为1个小时。
	DownloadUrl string `json:"download_url"`
}

func (o AuditionFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditionFile struct{}"
	}

	return strings.Join([]string{"AuditionFile", string(data)}, " ")
}
