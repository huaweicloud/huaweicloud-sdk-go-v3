package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateByPageRequest Request Object
type ShowTemplateByPageRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestTemplatePageParam `json:"body,omitempty"`
}

func (o ShowTemplateByPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateByPageRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateByPageRequest", string(data)}, " ")
}
