package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicKeyList struct {

	// 密钥列表
	Sshkey *[]PublicKey `json:"sshkey,omitempty" xml:"sshkey"`

	// 密钥总数
	Total *int32 `json:"total,omitempty" xml:"total"`
}

func (o PublicKeyList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicKeyList struct{}"
	}

	return strings.Join([]string{"PublicKeyList", string(data)}, " ")
}
