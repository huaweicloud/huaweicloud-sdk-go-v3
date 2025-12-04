package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceTopicReq struct {

	// **参数解释**： 待修改的Topic列表。 **约束限制**： 不涉及。
	Topics *[]UpdateInstanceTopicReqTopics `json:"topics,omitempty"`
}

func (o UpdateInstanceTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicReq struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicReq", string(data)}, " ")
}
