package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowTestCaseRequest Request Object
type BatchShowTestCaseRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestTestCasePageParam `json:"body,omitempty"`
}

func (o BatchShowTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowTestCaseRequest struct{}"
	}

	return strings.Join([]string{"BatchShowTestCaseRequest", string(data)}, " ")
}
