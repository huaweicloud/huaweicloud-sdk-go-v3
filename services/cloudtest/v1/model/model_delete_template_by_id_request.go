package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateByIdRequest Request Object
type DeleteTemplateByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 模板ID
	Id string `json:"id"`
}

func (o DeleteTemplateByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateByIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateByIdRequest", string(data)}, " ")
}
