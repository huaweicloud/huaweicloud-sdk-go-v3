package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobLabelResponse **参数解释**：作业标签，用于标识和分类作业。
type ShowSparkJobLabelResponse struct {

	// **参数解释**：标签键，用于标识标签的名称。 **取值范围**：长度为1~128个字符。
	Key string `json:"key"`

	// **参数解释**：标签值，用于标识标签的具体内容。 **取值范围**：长度为0~512个字符。
	Value string `json:"value"`
}

func (o ShowSparkJobLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobLabelResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobLabelResponse", string(data)}, " ")
}
