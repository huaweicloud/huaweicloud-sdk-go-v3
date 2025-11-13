package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssistantModelRequest Request Object
type UpdateAssistantModelRequest struct {

	// **参数解释**： 模型供应商ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	VendorId string `json:"vendor_id"`

	// **参数解释**： 模型ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	ModelId string `json:"model_id"`

	Body *UpdateAssistantModelReq `json:"body,omitempty"`
}

func (o UpdateAssistantModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssistantModelRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssistantModelRequest", string(data)}, " ")
}
