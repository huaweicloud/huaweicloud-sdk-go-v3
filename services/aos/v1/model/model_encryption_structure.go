package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 如果用户传递的var_value是已经加密过的，可以通过声名此项以要求资源编排服务在使用前进行解密，目前暂时只支持KMS加解密
type EncryptionStructure struct {
	Kms *KmsStructure `json:"kms"`
}

func (o EncryptionStructure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptionStructure struct{}"
	}

	return strings.Join([]string{"EncryptionStructure", string(data)}, " ")
}
