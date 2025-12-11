package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateConfigurationNameRequest Request Object
type ValidateConfigurationNameRequest struct {

	// **参数解释：** 参数模板名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name string `json:"name"`
}

func (o ValidateConfigurationNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConfigurationNameRequest struct{}"
	}

	return strings.Join([]string{"ValidateConfigurationNameRequest", string(data)}, " ")
}
