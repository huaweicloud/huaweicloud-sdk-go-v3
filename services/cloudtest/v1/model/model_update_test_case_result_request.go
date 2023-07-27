package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTestCaseResultRequest Request Object
type UpdateTestCaseResultRequest struct {

	// 项目唯一标识，固定长度32位字符
	ProjectId string `json:"project_id"`

	Body *UpdateTestCaseResultRequestBody `json:"body,omitempty"`
}

func (o UpdateTestCaseResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseResultRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseResultRequest", string(data)}, " ")
}
