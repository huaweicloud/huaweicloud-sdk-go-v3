package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlTdeInfo struct {

	// 是否打开透明加密功能。
	EnableTde bool `json:"enable_tde"`

	// 透明加密算法，支持AES256、SM4加密算法。
	EncryptionType string `json:"encryption_type"`
}

func (o MysqlTdeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlTdeInfo struct{}"
	}

	return strings.Join([]string{"MysqlTdeInfo", string(data)}, " ")
}
