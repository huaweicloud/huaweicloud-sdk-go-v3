package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppWhitelistPolicyHostResponseInfo 策略关联主机信息
type AppWhitelistPolicyHostResponseInfo struct {

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**: 主机发生事件数 **取值范围**: 最小值0，最大值2147483647
	EventNum *int32 `json:"event_num,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux。 - Windows：Windows。
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 学习状态 **约束限制**: 不涉及 **取值范围**: - effecting：学习完成，策略生效 - learned：学习完成，待确认 - learning：学习中 - pause：暂停 - abnormal：学习异常  **默认取值**: 不涉及
	LearningStatus *string `json:"learning_status,omitempty"`

	// **参数解释**： 是否应用 **取值范围**: - true：是 - false：否
	ApplyStatus *bool `json:"apply_status,omitempty"`

	// 是否开启阻断
	Intercept *bool `json:"intercept,omitempty"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略类型
	PolicyType *string `json:"policy_type,omitempty"`
}

func (o AppWhitelistPolicyHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppWhitelistPolicyHostResponseInfo struct{}"
	}

	return strings.Join([]string{"AppWhitelistPolicyHostResponseInfo", string(data)}, " ")
}
