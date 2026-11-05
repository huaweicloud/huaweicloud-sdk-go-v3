package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthGrantRequest struct {

	// **参数解释**：是否重置授权。 **约束限制**：不涉及。 **取值范围**： - true：重置已有委托（先删除再创建）。 - false：普通授权。 **默认取值**：false。
	IsReset *bool `json:"is_reset,omitempty"`
}

func (o AuthGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthGrantRequest struct{}"
	}

	return strings.Join([]string{"AuthGrantRequest", string(data)}, " ")
}
