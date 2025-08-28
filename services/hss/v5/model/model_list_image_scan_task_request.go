package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageScanTaskRequest Request Object
type ListImageScanTaskRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - local：本地镜像。 - registry：仓库镜像。  **默认取值**: 不涉及
	GlobalImageType *string `json:"global_image_type,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 任务总类型 **约束限制**: 不涉及 **取值范围**: - image_sync：镜像同步。 - image_scan：镜像扫描。  **默认取值**: 不涉及
	Type string `json:"type"`

	// **参数解释**: 任务细分类型 **约束限制**: 不涉及 **取值范围**: - cycle：定时扫描。 - manual：手动扫描。 - autoSync：定时同步。 - manualSync：手动同步。  **默认取值**: 不涉及
	TaskType *string `json:"task_type,omitempty"`

	// **参数解释**： 模糊匹配任务名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**： 任务id **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**： 任务创建时间，时间单位毫秒（ms） **约束限制**： 不涉及 **取值范围**： 最小值0，最大值4070880000000 **默认取值**： 不涉及
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 任务结束时间，时间单位毫秒（ms） **约束限制**： 不涉及 **取值范围**： 最小值0，最大值4070880000000 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**: 任务情况 **约束限制**: 不涉及 **取值范围**: - scanning：扫描中。 - finished：完成。  **默认取值**: 不涉及
	TaskStatus *string `json:"task_status,omitempty"`

	// **参数解释**: 扫描风险类型 **约束限制**: 不涉及 **取值范围**: - 0：none。 - 0x7fffffff：全部。 - 0x000f0000：漏洞。 - 0x0000f000：基线检查。 - 0x00000f00：恶意文件。 - 0x000000f0：敏感信息。 - 0x0000000f：软件合规。  **默认取值**: 不涉及
	ScanScope *int32 `json:"scan_scope,omitempty"`
}

func (o ListImageScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageScanTaskRequest struct{}"
	}

	return strings.Join([]string{"ListImageScanTaskRequest", string(data)}, " ")
}
