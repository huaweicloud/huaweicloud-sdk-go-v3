package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCloudStorageResponse Response Object
type BatchDeleteCloudStorageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteCloudStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCloudStorageResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteCloudStorageResponse", string(data)}, " ")
}
