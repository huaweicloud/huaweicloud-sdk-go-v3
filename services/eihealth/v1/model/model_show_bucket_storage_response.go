package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBucketStorageResponse Response Object
type ShowBucketStorageResponse struct {

	// 已用存量（字节）
	Size           *int64 `json:"size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowBucketStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBucketStorageResponse struct{}"
	}

	return strings.Join([]string{"ShowBucketStorageResponse", string(data)}, " ")
}
