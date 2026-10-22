package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiAppRequest Request Object
type UpdateAiAppRequest struct {

	// **参数解释：** AI应用ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AppId string `json:"app_id"`

	Body *AiAppUpdateRequest `json:"body,omitempty"`
}

func (o UpdateAiAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateAiAppRequest", string(data)}, " ")
}
