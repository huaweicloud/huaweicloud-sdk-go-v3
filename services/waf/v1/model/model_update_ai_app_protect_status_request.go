package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiAppProtectStatusRequest Request Object
type UpdateAiAppProtectStatusRequest struct {

	// **参数解释：** AI应用ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AppId string `json:"app_id"`

	Body *AiAppProtectStatusRequest `json:"body,omitempty"`
}

func (o UpdateAiAppProtectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiAppProtectStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateAiAppProtectStatusRequest", string(data)}, " ")
}
