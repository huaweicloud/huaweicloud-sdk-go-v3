package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestcaseByIdRequest Request Object
type ShowTestcaseByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 用例ID
	Id string `json:"id"`
}

func (o ShowTestcaseByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestcaseByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowTestcaseByIdRequest", string(data)}, " ")
}
