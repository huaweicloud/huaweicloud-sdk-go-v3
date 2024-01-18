package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PersistentStorage WKS存储。
type PersistentStorage struct {

	// WKS存储ID。
	Id *string `json:"id,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`

	StorageMetadata *StorageMetadata `json:"storage_metadata,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 个人目录声明数量。
	UserClaimCount *int32 `json:"user_claim_count,omitempty"`

	// 共享目录声明数量。
	ShareClaimCount *int32 `json:"share_claim_count,omitempty"`
}

func (o PersistentStorage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistentStorage struct{}"
	}

	return strings.Join([]string{"PersistentStorage", string(data)}, " ")
}
