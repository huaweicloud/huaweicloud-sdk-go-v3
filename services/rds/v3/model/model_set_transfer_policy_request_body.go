package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTransferPolicyRequestBody 转储策略请求体
type SetTransferPolicyRequestBody struct {

	// 自动转储策略
	Policy *[]TransferPolicyParam `json:"policy,omitempty"`
}

func (o SetTransferPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTransferPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetTransferPolicyRequestBody", string(data)}, " ")
}
