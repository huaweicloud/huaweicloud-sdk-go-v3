package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTestCasesRequest Request Object
type ListAllTestCasesRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *TestCasesQueryInfo `json:"body,omitempty"`
}

func (o ListAllTestCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTestCasesRequest struct{}"
	}

	return strings.Join([]string{"ListAllTestCasesRequest", string(data)}, " ")
}
