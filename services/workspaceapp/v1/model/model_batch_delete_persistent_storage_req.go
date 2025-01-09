package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePersistentStorageReq 根据storage_id删除对应的WKS存储。
type BatchDeletePersistentStorageReq struct {

	// storage_id,数量区间 [1, 50]。
	Items []string `json:"items"`
}

func (o BatchDeletePersistentStorageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePersistentStorageReq struct{}"
	}

	return strings.Join([]string{"BatchDeletePersistentStorageReq", string(data)}, " ")
}
