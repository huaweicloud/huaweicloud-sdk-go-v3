package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskTestCasesRequest Request Object
type ListTaskTestCasesRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *QueryTaskTestCasesInfo `json:"body,omitempty"`
}

func (o ListTaskTestCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskTestCasesRequest struct{}"
	}

	return strings.Join([]string{"ListTaskTestCasesRequest", string(data)}, " ")
}
