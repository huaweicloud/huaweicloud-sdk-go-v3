package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestCaseAndDefectInfoRequest Request Object
type ShowTestCaseAndDefectInfoRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *ShowTestCaseAndDefectInfoRequestBody `json:"body,omitempty"`
}

func (o ShowTestCaseAndDefectInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCaseAndDefectInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowTestCaseAndDefectInfoRequest", string(data)}, " ")
}
