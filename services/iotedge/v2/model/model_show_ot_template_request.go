package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowOtTemplateRequest struct {

	// 模板id，节点下唯一。
	TplId string `json:"tpl_id"`
}

func (o ShowOtTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOtTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowOtTemplateRequest", string(data)}, " ")
}
