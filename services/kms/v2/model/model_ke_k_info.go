package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 密钥详细信息。
type KeKInfo struct {

	// 密钥ID。
	KeyId *string `json:"key_id,omitempty" xml:"key_id"`

	// 用户域ID。
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`
}

func (o KeKInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeKInfo struct{}"
	}

	return strings.Join([]string{"KeKInfo", string(data)}, " ")
}
