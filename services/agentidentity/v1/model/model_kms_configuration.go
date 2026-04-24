package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KmsConfiguration struct {
	KeyType *KmsKeyType `json:"key_type"`

	// The identifier of the KMS key used for the token vault.
	KmsKeyId *string `json:"kms_key_id,omitempty"`
}

func (o KmsConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KmsConfiguration struct{}"
	}

	return strings.Join([]string{"KmsConfiguration", string(data)}, " ")
}
