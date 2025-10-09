package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityId struct {

	// **参数解释**： 节点上的容器安全组ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o SecurityId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityId struct{}"
	}

	return strings.Join([]string{"SecurityId", string(data)}, " ")
}
