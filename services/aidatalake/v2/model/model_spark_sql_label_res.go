package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkSqlLabelRes **参数解释**：作业标签信息，用于标识和分类作业。
type SparkSqlLabelRes struct {

	// **参数解释**：标签的键，用于标识标签的类型或类别。 **取值范围**：长度为1~128个字符。
	Key string `json:"key"`

	// **参数解释**：标签的值，用于标识标签的具体内容或属性。 **取值范围**：长度为1~128个字符。
	Value string `json:"value"`
}

func (o SparkSqlLabelRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlLabelRes struct{}"
	}

	return strings.Join([]string{"SparkSqlLabelRes", string(data)}, " ")
}
