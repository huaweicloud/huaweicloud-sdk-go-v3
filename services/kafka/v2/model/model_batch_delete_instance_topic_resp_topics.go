package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteInstanceTopicRespTopics struct {

	// **参数解释**： Topic名称。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 是否删除成功。 **取值范围**： - true：删除成功。 - false：删除失败。
	Success *bool `json:"success,omitempty"`
}

func (o BatchDeleteInstanceTopicRespTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTopicRespTopics struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTopicRespTopics", string(data)}, " ")
}
