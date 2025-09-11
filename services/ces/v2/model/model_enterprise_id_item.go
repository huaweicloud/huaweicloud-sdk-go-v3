package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterpriseIdItem struct {

	// **参数解释** 企业项目ID **约束限制** 不涉及 **取值范围** 只能包含小写字母、数字、“-”、“_”，可以自定义企业项目ID，长度为36个字符。也可以为0（代表默认企业项目ID）。 **默认取值** 不涉及
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

func (o EnterpriseIdItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseIdItem struct{}"
	}

	return strings.Join([]string{"EnterpriseIdItem", string(data)}, " ")
}
