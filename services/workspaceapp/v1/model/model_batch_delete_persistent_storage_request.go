package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePersistentStorageRequest Request Object
type BatchDeletePersistentStorageRequest struct {
	Body *BatchDeletePersistentStorageReq `json:"body,omitempty"`
}

func (o BatchDeletePersistentStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePersistentStorageRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePersistentStorageRequest", string(data)}, " ")
}
