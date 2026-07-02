package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartKafkaLogTaskReqEntity struct {

	// **参数解释**： 日志文件名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogFileName *string `json:"log_file_name,omitempty"`

	// **参数解释**： 日志组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogGroupName *string `json:"log_group_name,omitempty"`

	// **参数解释**： 日志流名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogStreamName *string `json:"log_stream_name,omitempty"`
}

func (o StartKafkaLogTaskReqEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartKafkaLogTaskReqEntity struct{}"
	}

	return strings.Join([]string{"StartKafkaLogTaskReqEntity", string(data)}, " ")
}
