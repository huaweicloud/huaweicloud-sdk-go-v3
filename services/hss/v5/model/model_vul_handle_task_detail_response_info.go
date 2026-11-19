package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulHandleTaskDetailResponseInfo 漏洞处置任务详情信息
type VulHandleTaskDetailResponseInfo struct {

	// **参数解释**: 主机ID **取值范围**: 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机名称 **取值范围**: 字符长度0-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 公网IP **取值范围**: 字符长度0-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 私网IP **取值范围**: 字符长度0-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 资产重要性 **取值范围**: - important：重要。 - common：一般。 - test：测试。
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 漏洞ID **取值范围**: 字符长度1-512位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞名称 **取值范围**: 字符长度0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 处置状态 **取值范围**: - ready：待执行。 - running：执行中。 - success：处置成功。 - failed：处置失败。
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 失败原因 **取值范围**: 字符长度0-512位
	FailedReason *string `json:"failed_reason,omitempty"`

	// **参数解释**: 失败原因的业务语言描述 **取值范围**: 字符长度0-512位
	FailedReasonDescription *string `json:"failed_reason_description,omitempty"`

	// **参数解释**: 失败原因的解决方案 **取值范围**: 字符长度0-512位
	FailedReasonSolution *string `json:"failed_reason_solution,omitempty"`

	// **参数解释**: 存在漏洞的软件名称 **取值范围**: 字符长度0-128位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 存在漏洞的软件版本 **取值范围**: 字符长度0-128位
	AppVersion *string `json:"app_version,omitempty"`
}

func (o VulHandleTaskDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHandleTaskDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"VulHandleTaskDetailResponseInfo", string(data)}, " ")
}
