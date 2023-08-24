package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppConfigsTemplateRequest Request Object
type ShowAppConfigsTemplateRequest struct {

	// 模板id，节点下唯一。
	TplId string `json:"tpl_id"`
}

func (o ShowAppConfigsTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppConfigsTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowAppConfigsTemplateRequest", string(data)}, " ")
}
