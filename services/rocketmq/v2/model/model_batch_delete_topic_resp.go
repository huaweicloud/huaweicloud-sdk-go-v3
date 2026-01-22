package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteTopicResp struct {

	// **参数解释**： 删除主题任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	JobId *string `json:"job_id,omitempty"`
}

func (o BatchDeleteTopicResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTopicResp struct{}"
	}

	return strings.Join([]string{"BatchDeleteTopicResp", string(data)}, " ")
}
