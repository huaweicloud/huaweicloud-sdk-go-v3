package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteTopicReq struct {

	// **参数解释**： 主题列表，当批量删除主题时使用。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topics *[]string `json:"topics,omitempty"`
}

func (o BatchDeleteTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTopicReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteTopicReq", string(data)}, " ")
}
