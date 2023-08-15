package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunTestCaseRequest Request Object
type RunTestCaseRequest struct {

	// 项目唯一标识，固定长度32位字符
	ProjectId string `json:"project_id"`

	Body *RunTestCaseRequestBody `json:"body,omitempty"`
}

func (o RunTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTestCaseRequest struct{}"
	}

	return strings.Join([]string{"RunTestCaseRequest", string(data)}, " ")
}
