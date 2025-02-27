package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportSecretsRequest struct {

	// 批量创建凭据参数
	Secrets []CreateSecretRequestBody `json:"secrets"`
}

func (o ImportSecretsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportSecretsRequest struct{}"
	}

	return strings.Join([]string{"ImportSecretsRequest", string(data)}, " ")
}
