package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTypeReq **参数解释**： 导出类型 **约束限制**: 不涉及 **取值范围**： 导出类型，取值为对应的查询接口的接口名。 - ListAntiVirusResult：病毒扫描结果列表导出。 - ListAntivirusHandleHistory：病毒扫描历史处置记录列表。 - ListAiInstances：AI实例记录列表。 - ListAiSessions：AI会话记录列表。 - ListAiTraceRiskDetail：AI行为轨迹风险详情列表。 - ListDefenderPlugins：智能体卫士插件列表。 - ListEventHandleHistory：告警事件历史处置记录列表。  **默认取值**: 不涉及
type ExportTypeReq struct {
}

func (o ExportTypeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTypeReq struct{}"
	}

	return strings.Join([]string{"ExportTypeReq", string(data)}, " ")
}
