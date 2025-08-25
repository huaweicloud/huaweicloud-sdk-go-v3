package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAppAccessKeyRequestBody struct {

	// 访问密钥名称
	KeyName string `json:"key_name"`

	// 访问密钥ak，默认为空，系统自动生成
	AccessKey *string `json:"access_key,omitempty"`

	// 访问密钥sk，默认为空，系统自动生成
	SecretKey *string `json:"secret_key,omitempty"`
}

func (o CreateAppAccessKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppAccessKeyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAppAccessKeyRequestBody", string(data)}, " ")
}
