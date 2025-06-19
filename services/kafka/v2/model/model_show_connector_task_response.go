package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectorTaskResponse Response Object
type ShowConnectorTaskResponse struct {

	// **参数解释**： Smart Connect任务名称。 **取值范围**： 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**： Smart Connect任务配置的Topic。 **取值范围**： 不涉及。
	Topics *string `json:"topics,omitempty"`

	// **参数解释**： Smart Connect任务配置的Topic正则表达式。 **取值范围**： 不涉及。
	TopicsRegex *string `json:"topics_regex,omitempty"`

	// **参数解释**： Smart Connect任务的源端类型。 **取值范围**： - NONE：不配置。 - KAFKA_REPLICATOR_SOURCE：Kafka数据复制。
	SourceType *string `json:"source_type,omitempty"`

	SourceTask *SmartConnectTaskRespSourceConfig `json:"source_task,omitempty"`

	// **参数解释**： Smart Connect任务的目标端类型。 **取值范围**： - NONE：不配置。 - OBS_SINK：转储。
	SinkType *string `json:"sink_type,omitempty"`

	SinkTask *SmartConnectTaskRespSinkConfig `json:"sink_task,omitempty"`

	// **参数解释**： Smart Connect任务的id。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： Smart Connect任务的状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： Smart Connect任务的创建时间。 **取值范围**： 不涉及。
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowConnectorTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectorTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectorTaskResponse", string(data)}, " ")
}
