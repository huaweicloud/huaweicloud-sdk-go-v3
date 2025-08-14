package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCloudStorageRequest Request Object
type BatchDeleteCloudStorageRequest struct {
	Body *BatchDeleteCloudStorageReq `json:"body,omitempty"`
}

func (o BatchDeleteCloudStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCloudStorageRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteCloudStorageRequest", string(data)}, " ")
}
