package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespDlq struct {

	// **参数解释**： 死信队列名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespDlq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespDlq struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespDlq", string(data)}, " ")
}
