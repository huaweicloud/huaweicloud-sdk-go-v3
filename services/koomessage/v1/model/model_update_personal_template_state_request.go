package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePersonalTemplateStateRequest struct {

	// 智能信息模板ID。
	TplId string `json:"tpl_id"`

	Body *UpdatePersonalTemplateStateDataRequest `json:"body,omitempty"`
}

func (o UpdatePersonalTemplateStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePersonalTemplateStateRequest struct{}"
	}

	return strings.Join([]string{"UpdatePersonalTemplateStateRequest", string(data)}, " ")
}
