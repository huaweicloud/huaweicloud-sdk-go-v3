package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyContent 策略内容
type PolicyContent struct {

	// 是否允许未扫描镜像启动
	EnableImageStartup bool `json:"enable_image_startup"`

	// 告警策略: - vuls: 漏洞 - baselines: 基线  - malwares: 恶意脚本
	PolicyType string `json:"policy_type"`

	// 风险等级
	Severity int32 `json:"severity"`

	// 危险项
	RiskyItem []string `json:"risky_item"`

	// 防护动作   - 0: 告警   - 1: 阻断   - 2: 放行
	Action int32 `json:"action"`
}

func (o PolicyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyContent struct{}"
	}

	return strings.Join([]string{"PolicyContent", string(data)}, " ")
}
