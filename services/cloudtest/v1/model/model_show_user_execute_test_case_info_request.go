package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserExecuteTestCaseInfoRequest Request Object
type ShowUserExecuteTestCaseInfoRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *ShowUserExecuteTestCaseInfoRequestBody `json:"body,omitempty"`
}

func (o ShowUserExecuteTestCaseInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserExecuteTestCaseInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowUserExecuteTestCaseInfoRequest", string(data)}, " ")
}
