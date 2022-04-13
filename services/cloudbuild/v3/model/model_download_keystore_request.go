package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DownloadKeystoreRequest struct {
	// 下载的文件名称

	FileName string `json:"file_name"`
	// 租户ID。32位数字、小写字母组合

	DomainId string `json:"domain_id"`
}

func (o DownloadKeystoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKeystoreRequest struct{}"
	}

	return strings.Join([]string{"DownloadKeystoreRequest", string(data)}, " ")
}
