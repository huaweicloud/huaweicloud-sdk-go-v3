package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOtTemplateRequest Request Object
type DeleteOtTemplateRequest struct {

	// 模板id，节点下唯一。
	TplId string `json:"tpl_id"`
}

func (o DeleteOtTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOtTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteOtTemplateRequest", string(data)}, " ")
}
