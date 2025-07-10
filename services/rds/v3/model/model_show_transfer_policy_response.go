package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransferPolicyResponse Response Object
type ShowTransferPolicyResponse struct {

	// 转储策略列表
	Policies       *[]TransferPolicy `json:"policies,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowTransferPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransferPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowTransferPolicyResponse", string(data)}, " ")
}
