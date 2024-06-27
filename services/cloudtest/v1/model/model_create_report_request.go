package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReportRequest Request Object
type CreateReportRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 版本uri
	VersionId string `json:"version_id"`

	Body *OprReportInfo `json:"body,omitempty"`
}

func (o CreateReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportRequest struct{}"
	}

	return strings.Join([]string{"CreateReportRequest", string(data)}, " ")
}
