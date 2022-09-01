package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateKmsTagsRequest struct {

	// 密钥ID
	KeyId string `json:"key_id" xml:"key_id"`

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	Body *BatchCreateKmsTagsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchCreateKmsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateKmsTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateKmsTagsRequest", string(data)}, " ")
}
