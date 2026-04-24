package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KmsKeyId The identifier of the KMS key used for the token vault.
type KmsKeyId struct {
}

func (o KmsKeyId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KmsKeyId struct{}"
	}

	return strings.Join([]string{"KmsKeyId", string(data)}, " ")
}
