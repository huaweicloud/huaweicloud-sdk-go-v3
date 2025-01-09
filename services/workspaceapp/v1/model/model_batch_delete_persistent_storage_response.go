package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePersistentStorageResponse Response Object
type BatchDeletePersistentStorageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeletePersistentStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePersistentStorageResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePersistentStorageResponse", string(data)}, " ")
}
