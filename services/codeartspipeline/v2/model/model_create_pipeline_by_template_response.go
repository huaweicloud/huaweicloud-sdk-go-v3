package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineByTemplateResponse Response Object
type CreatePipelineByTemplateResponse struct {

	// **参数解释**： 实例ID。 **取值范围**： 32位字符，由数字和字母组成。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePipelineByTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineByTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineByTemplateResponse", string(data)}, " ")
}
