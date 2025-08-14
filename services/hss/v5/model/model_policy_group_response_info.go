package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyGroupResponseInfo 策略组详情
type PolicyGroupResponseInfo struct {

	// **参数解释**: 策略组名称 **取值范围**: 字符长度1-256位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 策略组ID **取值范围**: 字符长度1-256位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 策略组描述 **取值范围**: 字符长度0-64位
	Description *string `json:"description,omitempty"`

	// **参数解释**: 关联服务器数 **取值范围**: 取值0-1000000
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**: 是否是默认策略组 **取值范围**: - true: 是默认策略组 - false: 不是默认策略组
	DefaultGroup *bool `json:"default_group,omitempty"`

	// **参数解释**: 是否可以删除，只有default_group为false且host_num为1时可以删除 **取值范围**: - true: 支持删除 - false: 不支持删除
	Deletable *bool `json:"deletable,omitempty"`

	// **参数解释**: 支持的操作系统 **取值范围**: - Linux: 支持Linux操作系统 - Windows: 支持Windows操作系统
	SupportOs *string `json:"support_os,omitempty"`

	// **参数解释**: 支持的版本 **取值范围**: - hss.version.advanced: 专业版 - hss.version.enterprise: 企业版 - hss.version.premium: 旗舰版 - hss.version.wtp: 网页防篡改版 - hss.version.container.enterprise: 容器版
	SupportVersion *string `json:"support_version,omitempty"`

	// **参数解释**: 防护模式 **取值范围**: - high_detection: 高检出模式 - equalization: 均衡模式
	ProtectMode *string `json:"protect_mode,omitempty"`
}

func (o PolicyGroupResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupResponseInfo struct{}"
	}

	return strings.Join([]string{"PolicyGroupResponseInfo", string(data)}, " ")
}
