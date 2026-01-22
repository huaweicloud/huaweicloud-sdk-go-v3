package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespInstanceIds struct {

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespInstanceIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespInstanceIds struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespInstanceIds", string(data)}, " ")
}
