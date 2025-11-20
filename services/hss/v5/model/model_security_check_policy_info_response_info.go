package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckPolicyInfoResponseInfo 配置检测策略信息
type SecurityCheckPolicyInfoResponseInfo struct {

	// **参数解释**: 策略详情 **取值范围**: 最小值0，最大值10241
	Content *string `json:"content,omitempty"`
}

func (o SecurityCheckPolicyInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckPolicyInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckPolicyInfoResponseInfo", string(data)}, " ")
}
