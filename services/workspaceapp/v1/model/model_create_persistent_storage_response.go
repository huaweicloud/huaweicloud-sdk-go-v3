package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePersistentStorageResponse Response Object
type CreatePersistentStorageResponse struct {

	// WKS存储ID
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	StorageMetadata *StorageMetadata `json:"storage_metadata,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 个人目录声明数量
	UserClaimCount *int32 `json:"user_claim_count,omitempty"`

	// 共享目录声明数量
	ShareClaimCount *int32 `json:"share_claim_count,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o CreatePersistentStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePersistentStorageResponse struct{}"
	}

	return strings.Join([]string{"CreatePersistentStorageResponse", string(data)}, " ")
}
