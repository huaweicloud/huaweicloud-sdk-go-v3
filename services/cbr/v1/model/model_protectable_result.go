package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectableResult struct {

	// 不支持备份的错误码
	Code *string `json:"code,omitempty"`

	// 不支持备份的原因
	Reason *string `json:"reason,omitempty"`

	// 是否可备份
	Result bool `json:"result"`

	Vault *VaultGet `json:"vault,omitempty"`

	// 资源不可备份的原因信息，当资源可保护性检验失败时才有该字段。
	Message *string `json:"message,omitempty"`
}

func (o ProtectableResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectableResult struct{}"
	}

	return strings.Join([]string{"ProtectableResult", string(data)}, " ")
}
