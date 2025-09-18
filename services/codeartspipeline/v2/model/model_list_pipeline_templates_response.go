package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineTemplatesResponse Response Object
type ListPipelineTemplatesResponse struct {

	// **参数解释**： 起始偏移。 **取值范围**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 查询大小。 **取值范围**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 记录总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 流水线模板列表，包含流水线模板的详细信息。 **取值范围**： 不涉及。
	Templates      *[]PipelineTemplateSimpleVo `json:"templates,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListPipelineTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListPipelineTemplatesResponse", string(data)}, " ")
}
