package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicKeyList struct {
	// 密钥列表

	Sshkey *[]PublicKey `json:"sshkey,omitempty"`
	// 密钥总数

	Total *int32 `json:"total,omitempty"`
}

func (o PublicKeyList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicKeyList struct{}"
	}

	return strings.Join([]string{"PublicKeyList", string(data)}, " ")
}
