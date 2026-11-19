package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCommonInterfaceRequestBody 公共的字段信息
type ExportCommonInterfaceRequestBody struct {

	// **参数解释**： 导出类型 **约束限制**: 不涉及 **取值范围**： 导出类型，取值为对应的查询接口的接口名。 - ListAntiVirusResult：病毒扫描结果列表导出。 - ListAntivirusHandleHistory：病毒扫描历史处置记录列表。 - ListAiInstances：AI实例记录列表。 - ListAiSessions：AI会话记录列表。 - ListAiTraceRiskDetail：AI行为轨迹风险详情列表。 - ListDefenderPlugins：智能体卫士插件列表。 - ListEventHandleHistory：告警事件历史处置记录列表。  **默认取值**: 不涉及
	ExportType string `json:"export_type"`

	// **参数解释**: 查询条件信息，把对应的json转化成字符串，举例如下：当export_type=ListAntiVirusResult，传递查询条件为：{\"severities\":\"Critical\",\"handle_status\":\"unhandled\",\"asset_value\":\"important\"}，查询条件信息可以通过对应的查询接口获取。 **约束限制**: 不涉及 **取值范围**: 字符长度1-5000位 **默认取值**: 不涉及
	Condition *string `json:"condition,omitempty"`

	// 导出的数据的Id列表, 当id_list有值时，会忽略condition的信息: - export_type = ListAiInstances时对应的是agent_id - export_type = ListAiSessions时对应的是session_id - export_type = ListAiTraceRiskDetail时对应的是规则详情记录id - export_type = ListDefenderPlugins时对应的是instance_id
	IdList *[]string `json:"id_list,omitempty"`
}

func (o ExportCommonInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCommonInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"ExportCommonInterfaceRequestBody", string(data)}, " ")
}
