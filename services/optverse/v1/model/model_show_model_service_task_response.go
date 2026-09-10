package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModelServiceTaskResponse Response Object
type ShowModelServiceTaskResponse struct {

	// **参数解释**： 任务ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 任务类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 任务状态。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： 错误信息。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**： 创建时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**： 更新时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 任务输出。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Outputs        map[string]interface{} `json:"outputs,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowModelServiceTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModelServiceTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowModelServiceTaskResponse", string(data)}, " ")
}
