package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppWhitelistPolicyProcessRequest Request Object
type ListAppWhitelistPolicyProcessRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**： 策略ID **约束限制**： 必填 **取值范围**： 字符长度1-64位 **默认取值**： 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释**： 信任状态 **约束限制**: 不涉及 **取值范围**: - trust：可信 - suspicious：可疑 - malicious：未知 - unknown：未知  **默认取值**: 不涉及
	ProcessStatus *string `json:"process_status,omitempty"`

	// **参数解释**： 进程类型 **约束限制**: 不涉及 **取值范围**: - system：系统程序 - interpretive：解释类程序 - normal：普通可执行程序  **默认取值**: 不涉及
	ProcessType *string `json:"process_type,omitempty"`

	// **参数解释**： 进程名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
	ProcessName *string `json:"process_name,omitempty"`

	// **参数解释**： 进程hash **约束限制**： 不涉及 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
	ProcessHash *string `json:"process_hash,omitempty"`

	// **参数解释**： 进程路径 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	ProcessPath *string `json:"process_path,omitempty"`

	// **参数解释**： 确认状态 **约束限制**: 不涉及 **取值范围**: - confirmed：已确认 - unconfirmed：未确认  **默认取值**: 不涉及
	HandleStatus *string `json:"handle_status,omitempty"`

	// 操作系统类型，包含如下2种。   - Linux：Linux。   - Windows：Windows。
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 文件签名 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
	FileSigner *string `json:"file_signer,omitempty"`
}

func (o ListAppWhitelistPolicyProcessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppWhitelistPolicyProcessRequest struct{}"
	}

	return strings.Join([]string{"ListAppWhitelistPolicyProcessRequest", string(data)}, " ")
}
