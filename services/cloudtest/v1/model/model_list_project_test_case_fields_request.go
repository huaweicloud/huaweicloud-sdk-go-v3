package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTestCaseFieldsRequest Request Object
type ListProjectTestCaseFieldsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ListProjectTestCaseFieldsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTestCaseFieldsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTestCaseFieldsRequest", string(data)}, " ")
}
