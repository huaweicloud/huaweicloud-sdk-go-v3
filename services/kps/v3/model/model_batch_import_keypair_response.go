package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchImportKeypairResponse Response Object
type BatchImportKeypairResponse struct {

	// 导入失败的SSH密钥对信息及导入失败的原因
	FailedKeypairs *[]FailedKeypair `json:"failed_keypairs,omitempty"`

	// 成功导入的SSH密钥对信息
	SucceededKeypairs *[]CreateKeypairResponseBody `json:"succeeded_keypairs,omitempty"`
	HttpStatusCode    int                          `json:"-"`
}

func (o BatchImportKeypairResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportKeypairResponse struct{}"
	}

	return strings.Join([]string{"BatchImportKeypairResponse", string(data)}, " ")
}
