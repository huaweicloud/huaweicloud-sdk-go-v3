package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAimPersonalTemplateResponseMode struct {

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`
}

func (o CreateAimPersonalTemplateResponseMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAimPersonalTemplateResponseMode struct{}"
	}

	return strings.Join([]string{"CreateAimPersonalTemplateResponseMode", string(data)}, " ")
}
