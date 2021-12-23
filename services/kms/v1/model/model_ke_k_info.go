package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeKInfo struct {
	// 密钥ID。

	KeyId *string `json:"key_id,omitempty"`
	// 用户域ID。

	DomainId *string `json:"domain_id,omitempty"`
}

func (o KeKInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeKInfo struct{}"
	}

	return strings.Join([]string{"KeKInfo", string(data)}, " ")
}
