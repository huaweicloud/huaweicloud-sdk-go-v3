package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssistantModelResponse Response Object
type CreateAssistantModelResponse struct {

	// **参数解释**： 助手模型ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssistantModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssistantModelResponse struct{}"
	}

	return strings.Join([]string{"CreateAssistantModelResponse", string(data)}, " ")
}
