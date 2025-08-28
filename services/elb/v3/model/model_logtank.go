package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Logtank struct {

	// **参数解释**：云日志ID。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：项目ID。获取方式请参见[获取项目ID](elb_fl_0008.xml)。  **取值范围**：长度为32个字符，由小写字母和数字组成。
	ProjectId string `json:"project_id"`

	// **参数解释**：负载均衡器ID。  **取值范围**：不涉及
	LoadbalancerId string `json:"loadbalancer_id"`

	// **参数解释**：云日志分组ID。  **取值范围**：不涉及
	LogGroupId string `json:"log_group_id"`

	// **参数解释**：云日志主题ID。  **取值范围**：不涉及
	LogTopicId string `json:"log_topic_id"`
}

func (o Logtank) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Logtank struct{}"
	}

	return strings.Join([]string{"Logtank", string(data)}, " ")
}
