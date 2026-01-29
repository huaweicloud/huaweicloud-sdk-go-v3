package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceConfigurationsRequestBody struct {

	// **参数解释：** 参数值对象，用户基于默认参数模板自定义的参数值。 **约束限制:** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 为空时不修改参数值。
	Values map[string]string `json:"values"`
}

func (o UpdateInstanceConfigurationsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationsRequestBody", string(data)}, " ")
}
