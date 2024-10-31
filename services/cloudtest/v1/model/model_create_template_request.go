package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateRequest Request Object
type CreateTemplateRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestSaveTemplateParam `json:"body,omitempty"`
}

func (o CreateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateRequest", string(data)}, " ")
}
