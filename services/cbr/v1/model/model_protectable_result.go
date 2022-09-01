package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectableResult struct {

	// 不支持备份的错误码
	Code *string `json:"code,omitempty" xml:"code"`

	// 不支持备份的原因
	Reason *string `json:"reason,omitempty" xml:"reason"`

	// 是否可备份
	Result bool `json:"result" xml:"result"`

	Vault *VaultGet `json:"vault,omitempty" xml:"vault"`

	// 资源不可备份的原因信息，当资源可保护性检验失败时才有该字段。
	Message *string `json:"message,omitempty" xml:"message"`
}

func (o ProtectableResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectableResult struct{}"
	}

	return strings.Join([]string{"ProtectableResult", string(data)}, " ")
}
