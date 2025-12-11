package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntiVirusResultRequest Request Object
type ListAntiVirusResultRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器弹性IP地址 **约束限制**: 不涉及 **取值范围**: IPv4格式（长度7-15位）、IPv6格式（长度15-39位） **默认取值**: 无
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 处置状态 **约束限制**: 不涉及 **取值范围**: 包含如下:   - unhandled：未处理   - handled：已处理 **默认取值**: 不涉及
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 威胁等级 **约束限制**: 不涉及 **取值范围**: 威胁等级，包含如下:   - Low：低危   - Medium：中危   - High：高危   - Critical：致命 **默认取值**: 不涉及
	SeverityList *[]string `json:"severity_list,omitempty"`

	// **参数解释**： 资产重要性 **约束限制**： 不涉及 **取值范围**： - important：重要资产 - common：一般资产 - test：测试资产  **默认取值**： 无
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 病毒名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	MalwareName *string `json:"malware_name,omitempty"`

	// **参数解释**： 文件路径 **约束限制**： 不涉及 **取值范围**： 字符数1-512位 **默认取值**： 不涉及
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**： 文件hash，当前为sha256 **约束限制**： 需为合法的SHA256哈希值 **取值范围**： 字符长度64位 **默认取值**： 不涉及
	FileHash *string `json:"file_hash,omitempty"`

	// **参数解释**: 任务名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**: 是否使用手动隔离按钮 **约束限制**: 不涉及 **取值范围**: true（已使用手动隔离）、false（未使用手动隔离） **默认取值**: 不涉及
	ManualIsolate *bool `json:"manual_isolate,omitempty"`
}

func (o ListAntiVirusResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntiVirusResultRequest struct{}"
	}

	return strings.Join([]string{"ListAntiVirusResultRequest", string(data)}, " ")
}
