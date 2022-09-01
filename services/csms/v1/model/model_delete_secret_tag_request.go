package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSecretTagRequest struct {

	// 凭据ID
	SecretId string `json:"secret_id" xml:"secret_id"`

	// 标签键的值
	Key string `json:"key" xml:"key"`
}

func (o DeleteSecretTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecretTagRequest", string(data)}, " ")
}
