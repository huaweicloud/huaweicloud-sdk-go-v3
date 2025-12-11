package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountResponseInfo 事件列表详情
type AccountResponseInfo struct {

	// **参数解释**: 账号的唯一名称，用于标识账号身份； **取值范围**: 字符长度1-64位，支持字母、数字、连字符、下划线，不能以特殊字符开头或结尾
	AccountName *string `json:"account_name,omitempty"`

	// **参数解释**: 账号的唯一标识ID，用于唯一确定某个账号； **取值范围**: 字符长度1-64位，符合平台账号ID命名规范（如UUID或数字组合）
	AccountId *string `json:"account_id,omitempty"`

	// **参数解释**: 账号所属组织的唯一标识ID； **取值范围**: 字符长度1-64位，符合平台组织ID命名规范
	OrganizationId *string `json:"organization_id,omitempty"`

	// **参数解释**: 账号所属项目的唯一标识ID； **取值范围**: 字符长度1-64位，符合平台项目ID命名规范；
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 账号所属项目的名称，用于直观标识项目； **取值范围**: 字符长度1-64位，支持字母、数字、连字符、下划线及中文，无复杂度额外要求
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释**: 当前账号下已关联的主机总数量； **取值范围**: 非负整数，最小值0，最大值取决于平台资源配额；单位：台
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**: 当前账号下主机存在的漏洞风险总数量； **取值范围**: 非负整数，取值范围0-2147483647；单位：个
	VulnerabilityNum *int32 `json:"vulnerability_num,omitempty"`

	// **参数解释**: 当前账号下主机基线检测未通过的风险总数量； **取值范围**: 非负整数，取值范围0-2147483647；单位：个
	BaselineNum *int32 `json:"baseline_num,omitempty"`

	// **参数解释**: 当前账号下主机发生的安全入侵告警总数量； **取值范围**: 非负整数，取值范围0-2147483647；单位：条
	IntrusionNum *int32 `json:"intrusion_num,omitempty"`
}

func (o AccountResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountResponseInfo struct{}"
	}

	return strings.Join([]string{"AccountResponseInfo", string(data)}, " ")
}
