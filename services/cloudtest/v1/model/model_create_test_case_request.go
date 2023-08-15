package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTestCaseRequest Request Object
type CreateTestCaseRequest struct {

	// 项目唯一标识，固定长度32位字符
	ProjectId string `json:"project_id"`

	Body *CreateTestCaseRequestBody `json:"body,omitempty"`
}

func (o CreateTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestCaseRequest struct{}"
	}

	return strings.Join([]string{"CreateTestCaseRequest", string(data)}, " ")
}
