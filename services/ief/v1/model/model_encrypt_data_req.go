package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EncryptDataReq 加密数据配置
type EncryptDataReq struct {
	EncryptData *EncryptDataIn `json:"encrypt_data"`
}

func (o EncryptDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataReq struct{}"
	}

	return strings.Join([]string{"EncryptDataReq", string(data)}, " ")
}
