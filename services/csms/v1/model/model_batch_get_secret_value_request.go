package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchGetSecretValueRequest Request Object
type BatchGetSecretValueRequest struct {
	Body *BatchGetSecretValueRequestBody `json:"body,omitempty"`
}

func (o BatchGetSecretValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchGetSecretValueRequest struct{}"
	}

	return strings.Join([]string{"BatchGetSecretValueRequest", string(data)}, " ")
}
