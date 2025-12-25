package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadIndicatorTemplateRequest Request Object
type DownloadIndicatorTemplateRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o DownloadIndicatorTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadIndicatorTemplateRequest struct{}"
	}

	return strings.Join([]string{"DownloadIndicatorTemplateRequest", string(data)}, " ")
}
