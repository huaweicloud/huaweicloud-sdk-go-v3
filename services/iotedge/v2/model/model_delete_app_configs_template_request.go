package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppConfigsTemplateRequest Request Object
type DeleteAppConfigsTemplateRequest struct {

	// 模板id，节点下唯一。
	TplId string `json:"tpl_id"`
}

func (o DeleteAppConfigsTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppConfigsTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppConfigsTemplateRequest", string(data)}, " ")
}
