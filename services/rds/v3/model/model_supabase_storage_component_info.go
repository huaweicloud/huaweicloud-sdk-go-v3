package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseStorageComponentInfo Storage组件信息（show返回ak，不返回sk）。
type SupabaseStorageComponentInfo struct {
	Enabled *bool `json:"enabled,omitempty"`

	Type *string `json:"type,omitempty"`

	BucketName *string `json:"bucket_name,omitempty"`

	Ak *string `json:"ak,omitempty"`
}

func (o SupabaseStorageComponentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseStorageComponentInfo struct{}"
	}

	return strings.Join([]string{"SupabaseStorageComponentInfo", string(data)}, " ")
}
