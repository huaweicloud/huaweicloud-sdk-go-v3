package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateGroupResp struct {

	// **参数解释**： 创建成功的消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o CreateGroupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupResp struct{}"
	}

	return strings.Join([]string{"CreateGroupResp", string(data)}, " ")
}
