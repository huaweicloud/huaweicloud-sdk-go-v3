package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePersonalTemplateStateDataRequest 启用禁用模板请求体。
type UpdatePersonalTemplateStateDataRequest struct {

	// 状态。 - 0：禁用 - 1：启用
	TplState *int32 `json:"tpl_state,omitempty"`
}

func (o UpdatePersonalTemplateStateDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePersonalTemplateStateDataRequest struct{}"
	}

	return strings.Join([]string{"UpdatePersonalTemplateStateDataRequest", string(data)}, " ")
}
