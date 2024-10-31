package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultTemplateByPageRequest Request Object
type ShowDefaultTemplateByPageRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestGetDefaultTemplateParam `json:"body,omitempty"`
}

func (o ShowDefaultTemplateByPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultTemplateByPageRequest struct{}"
	}

	return strings.Join([]string{"ShowDefaultTemplateByPageRequest", string(data)}, " ")
}
