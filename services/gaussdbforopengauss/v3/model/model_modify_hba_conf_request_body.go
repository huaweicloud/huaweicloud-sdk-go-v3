package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyHbaConfRequestBody struct {

	// **参数解释**: 需要修改的hba配置信息。 **约束限制**: 不涉及。
	BeforeConfs *interface{} `json:"before_confs"`

	// **参数解释**: 修改后的hba配置信息。 **约束限制**: 不涉及。
	AfterConfs *interface{} `json:"after_confs"`
}

func (o ModifyHbaConfRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHbaConfRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyHbaConfRequestBody", string(data)}, " ")
}
