package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTestReportRequest Request Object
type BatchDeleteTestReportRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *[]string `json:"body,omitempty"`
}

func (o BatchDeleteTestReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTestReportRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestReportRequest", string(data)}, " ")
}
