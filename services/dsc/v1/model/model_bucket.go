package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Bucket struct {

	// 资产名称
	AssetName *string `json:"asset_name,omitempty"`

	// 桶位置
	BucketLocation *string `json:"bucket_location,omitempty"`

	// 桶名称
	BucketName *string `json:"bucket_name,omitempty"`

	// 桶策略
	BucketPolicy *string `json:"bucket_policy,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 是否被删除
	Deleted *bool `json:"deleted,omitempty"`

	// 桶ID
	Id *string `json:"id,omitempty"`

	// 是否被删除
	IsDeleted *bool `json:"is_deleted,omitempty"`
}

func (o Bucket) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Bucket struct{}"
	}

	return strings.Join([]string{"Bucket", string(data)}, " ")
}
