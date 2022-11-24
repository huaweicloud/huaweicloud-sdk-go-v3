package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSecretTagsRequest struct {

	// 凭据ID
	SecretId string `json:"secret_id"`
}

func (o ListSecretTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSecretTagsRequest", string(data)}, " ")
}
