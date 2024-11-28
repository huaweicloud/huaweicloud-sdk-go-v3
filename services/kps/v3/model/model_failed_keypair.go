package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailedKeypair struct {

	// SSH密钥对名称
	KeypairName string `json:"keypair_name"`

	// 导入失败的原因
	FailedReason string `json:"failed_reason"`
}

func (o FailedKeypair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedKeypair struct{}"
	}

	return strings.Join([]string{"FailedKeypair", string(data)}, " ")
}
