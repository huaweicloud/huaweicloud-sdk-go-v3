package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineTemplateResponse Response Object
type CreatePipelineTemplateResponse struct {

	// **参数解释**： 模板ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TemplateId     *string `json:"templateId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePipelineTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineTemplateResponse", string(data)}, " ")
}
