package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAimPersonalTemplateRequest Request Object
type DeleteAimPersonalTemplateRequest struct {

	// 智能信息模板ID。
	TplId string `json:"tpl_id"`
}

func (o DeleteAimPersonalTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAimPersonalTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteAimPersonalTemplateRequest", string(data)}, " ")
}
