package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HtapErrorLogDetailResponseErrorLogList struct {

	// **参数解释**： HTAP标准版实例节点ID。  **取值范围**：  不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**： 日志时间。  **取值范围**：  不涉及。
	Time string `json:"time"`

	// **参数解释**： 日志级别。  **取值范围**：  不涉及。
	Level string `json:"level"`

	// **参数解释**： 日志内容。  **取值范围**：  不涉及。
	Content string `json:"content"`

	// **参数解释**： 日志单行序列号，第一次查询时不需要此参数，后续分页查询时需要使用，可从上次查询的返回信息中获取。  **取值范围**：  不涉及。
	LineNum string `json:"line_num"`
}

func (o HtapErrorLogDetailResponseErrorLogList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapErrorLogDetailResponseErrorLogList struct{}"
	}

	return strings.Join([]string{"HtapErrorLogDetailResponseErrorLogList", string(data)}, " ")
}
