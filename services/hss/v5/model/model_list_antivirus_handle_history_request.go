package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntivirusHandleHistoryRequest Request Object
type ListAntivirusHandleHistoryRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 病毒名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	MalwareName *string `json:"malware_name,omitempty"`

	// **参数解释**： 文件路径 **约束限制**： 不涉及 **取值范围**： 字符数1-512位 **默认取值**： 不涉及
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**: 威胁等级 **约束限制**: 不涉及 **取值范围**: 威胁等级，包含如下:   - Low：低危   - Medium：中危   - High：高危   - Critical：致命 **默认取值**: 不涉及
	SeverityList *[]string `json:"severity_list,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器弹性IP地址 **约束限制**: 不涉及 **取值范围**: IPv4格式（长度7-15位）、IPv6格式（长度15-39位） **默认取值**: 无
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 资产重要性 **约束限制**： 不涉及 **取值范围**： - important：重要资产 - common：一般资产 - test：测试资产  **默认取值**： 无
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 处理方式 **约束限制**: 不涉及 **取值范围**: 处理方式，包含如下:   - mark_as_handled：手动处理   - ignore：忽略   - add_to_alarm_whitelist：加入告警白名单   - isolate_and_kill：隔离文件   - unhandle：取消手动处理   - do_not_ignore：取消忽略   - remove_from_alarm_whitelist：删除告警白名单   - do_not_isolate_or_kill：取消隔离文件 **默认取值**: 不涉及
	HandleMethod *string `json:"handle_method,omitempty"`

	// **参数解释**: 用户名 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 排序的顺序 **约束限制**: 不涉及 **取值范围**:   - asc  : 正序   - desc : 倒序  **默认取值**: 正序排序
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**: 事件类型 **约束限制**: 不涉及 **取值范围**: 0（病毒查杀事件）、1（恶意文件处置事件） **默认取值**: 不涉及
	EventType *int32 `json:"event_type,omitempty"`

	// **参数解释**: 排序字段 **约束限制**: 不涉及 **取值范围**: handle_time（处置时间） **默认取值**: 不涉及
	SortKey *string `json:"sort_key,omitempty"`
}

func (o ListAntivirusHandleHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntivirusHandleHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListAntivirusHandleHistoryRequest", string(data)}, " ")
}
