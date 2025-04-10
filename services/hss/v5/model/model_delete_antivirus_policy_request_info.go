package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAntivirusPolicyRequestInfo 删除自定义扫描策略
type DeleteAntivirusPolicyRequestInfo struct {

	// 策略Id列表
	PolicyIds *[]string `json:"policy_ids,omitempty"`
}

func (o DeleteAntivirusPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAntivirusPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"DeleteAntivirusPolicyRequestInfo", string(data)}, " ")
}
