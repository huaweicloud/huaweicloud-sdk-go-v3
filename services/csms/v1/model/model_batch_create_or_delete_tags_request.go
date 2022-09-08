package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateOrDeleteTagsRequest struct {

	// 凭据ID
	SecretId string `json:"secret_id"`

	Body *BatchCreateOrDeleteTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteTagsRequest", string(data)}, " ")
}
