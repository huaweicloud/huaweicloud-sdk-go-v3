package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudStorageRequest Request Object
type DeleteCloudStorageRequest struct {

	// WKS存储ID。
	StorageId string `json:"storage_id"`
}

func (o DeleteCloudStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudStorageRequest struct{}"
	}

	return strings.Join([]string{"DeleteCloudStorageRequest", string(data)}, " ")
}
