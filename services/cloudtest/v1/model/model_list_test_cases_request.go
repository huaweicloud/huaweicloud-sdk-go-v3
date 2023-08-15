package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCasesRequest Request Object
type ListTestCasesRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *ListTestCasesRequestBody `json:"body,omitempty"`
}

func (o ListTestCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCasesRequest struct{}"
	}

	return strings.Join([]string{"ListTestCasesRequest", string(data)}, " ")
}
