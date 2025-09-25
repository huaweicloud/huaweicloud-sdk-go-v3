package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVaultResponse Response Object
type CreateVaultResponse struct {
	Vault *VaultCreateResource `json:"vault,omitempty"`

	// 包周期创建订单信息（仅包周期创建时显示）
	Orders *[]CbcOrderResult `json:"orders,omitempty"`

	// 包周期订购结果（仅包周期创建时显示）
	RetCode *int32 `json:"retCode,omitempty"`

	// 包周期创建错误信息（仅包周期创建时显示）
	ErrText *string `json:"errText,omitempty"`

	// 包周期创建错误码（仅包周期创建时显示）
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVaultResponse struct{}"
	}

	return strings.Join([]string{"CreateVaultResponse", string(data)}, " ")
}
