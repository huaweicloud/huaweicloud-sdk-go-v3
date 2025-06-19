package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKeystoreRequest Request Object
type DeleteKeystoreRequest struct {

	// 文件秘钥Id
	KeystoreId string `json:"keystore_id"`
}

func (o DeleteKeystoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeystoreRequest struct{}"
	}

	return strings.Join([]string{"DeleteKeystoreRequest", string(data)}, " ")
}
