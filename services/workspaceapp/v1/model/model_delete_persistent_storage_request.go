package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePersistentStorageRequest Request Object
type DeletePersistentStorageRequest struct {

	// WKS存储ID。
	StorageId string `json:"storage_id"`
}

func (o DeletePersistentStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePersistentStorageRequest struct{}"
	}

	return strings.Join([]string{"DeletePersistentStorageRequest", string(data)}, " ")
}
