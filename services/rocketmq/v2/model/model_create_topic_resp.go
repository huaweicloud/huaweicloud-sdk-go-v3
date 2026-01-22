package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTopicResp struct {

	// **参数解释**： 主题名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`
}

func (o CreateTopicResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicResp struct{}"
	}

	return strings.Join([]string{"CreateTopicResp", string(data)}, " ")
}
