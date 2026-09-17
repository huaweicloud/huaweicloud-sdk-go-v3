package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineTagResp **参数解释**： 标签响应体。 **取值范围**： 不涉及。
type PipelineTagResp struct {

	// **参数解释**： 标签ID。 **取值范围**： 32位字符，由数字和字母组成。
	TagId *string `json:"tag_id,omitempty"`

	// **参数解释**： 标签名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 标签颜色。 **取值范围**： 不涉及。
	Color *string `json:"color,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 32位字符，由数字和字母组成。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 项目名称。 **取值范围**： 不涉及。
	ProjectName *string `json:"project_name,omitempty"`
}

func (o PipelineTagResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTagResp struct{}"
	}

	return strings.Join([]string{"PipelineTagResp", string(data)}, " ")
}
