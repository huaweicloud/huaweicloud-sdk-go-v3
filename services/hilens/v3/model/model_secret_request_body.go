package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecretRequestBody 创建密钥或者更新密钥的请求体
type SecretRequestBody struct {

	// 工作空间ID。默认为default。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Secret *SecretObjectInSecretRequestBody `json:"secret"`
}

func (o SecretRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecretRequestBody struct{}"
	}

	return strings.Join([]string{"SecretRequestBody", string(data)}, " ")
}
