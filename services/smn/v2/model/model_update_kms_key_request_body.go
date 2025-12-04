package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateKmsKeyRequestBody struct {

	// 在DEW服务上创建的用户主密钥对应的密钥ID，具体参考在DEW服务上创建密钥章节。
	KeyId string `json:"key_id"`
}

func (o UpdateKmsKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKmsKeyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKmsKeyRequestBody", string(data)}, " ")
}
