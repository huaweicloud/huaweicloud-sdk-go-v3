package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopNamePolicyReq 删除桌面名称策略请求。
type BatchDeleteDesktopNamePolicyReq struct {

	// 策略id数组，最多支持50。
	PolicyIds *[]string `json:"policy_ids,omitempty"`
}

func (o BatchDeleteDesktopNamePolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopNamePolicyReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopNamePolicyReq", string(data)}, " ")
}
