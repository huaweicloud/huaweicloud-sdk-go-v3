package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Context struct {

	// **参数解释**： 上下文cluster信息 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Cluster *string `json:"cluster,omitempty"`

	// **参数解释**： 上下文user信息 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	User *string `json:"user,omitempty"`
}

func (o Context) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Context struct{}"
	}

	return strings.Join([]string{"Context", string(data)}, " ")
}
