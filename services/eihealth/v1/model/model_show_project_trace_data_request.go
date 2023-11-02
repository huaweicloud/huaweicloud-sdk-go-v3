package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectTraceDataRequest Request Object
type ShowProjectTraceDataRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 审计日志全路径（路径）
	Path string `json:"path"`
}

func (o ShowProjectTraceDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectTraceDataRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectTraceDataRequest", string(data)}, " ")
}
