package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoLaunchChangeHistoriesRequest Request Object
type ListAutoLaunchChangeHistoriesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 服务器的唯一标识ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器IP **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 自启动项名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	AutoLaunchName *string `json:"auto_launch_name,omitempty"`

	// **参数解释**: 自启动项类型 **约束限制**: 不涉及 **取值范围**: - 0：自启动服务 - 1：定时任务 - 2：预加载动态库 - 3：Run注册表键 - 4：开机启动文件夹  **默认取值**: 不涉及
	Type *int32 `json:"type,omitempty"`

	// **参数解释**: 变更类型 **约束限制**: 不涉及 **取值范围**: - add：新建 - delete：删除 - modify：修改  **默认取值**: 不涉及
	VariationType *string `json:"variation_type,omitempty"`

	// **参数解释**: 排序的key值，目前只支持按照recent_scan_time排序，按照recent_scan_time排序时，根据sort_dir的值决定升序还是降序 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 排序方式，默认为降序，当sort_key为按照recent_scan_time排序时，根据当前值决定升序还是降序，当sort_key为其他值时均为降序 **约束限制**: 不涉及 **取值范围**: - asc：升序 - desc：降序  **默认取值**: 不涉及
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 开始时间，13位时间戳 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**: 结束时间，13位时间戳 **约束限制**: 需大于等于begin_time，未传begin_time时默认从时间戳0开始查询 **取值范围**: 最小值0，最大值9223372036854775807（UTC时区，从1970-01-01 00:00:00开始计算） **默认取值**: 不涉及
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ListAutoLaunchChangeHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchChangeHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchChangeHistoriesRequest", string(data)}, " ")
}
