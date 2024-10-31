package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAssetTemplateRequest Request Object
type DownloadAssetTemplateRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o DownloadAssetTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAssetTemplateRequest struct{}"
	}

	return strings.Join([]string{"DownloadAssetTemplateRequest", string(data)}, " ")
}
