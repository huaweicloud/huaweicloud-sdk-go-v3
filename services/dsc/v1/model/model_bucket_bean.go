package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BucketBean struct {

	// 资产名称
	AssetName *string `json:"asset_name,omitempty"`

	// 桶位置
	Location *string `json:"location,omitempty"`

	// 桶名称
	BucketName *string `json:"bucket_name,omitempty"`

	// 桶策略
	BucketPolicy *string `json:"bucket_policy,omitempty"`
}

func (o BucketBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BucketBean struct{}"
	}

	return strings.Join([]string{"BucketBean", string(data)}, " ")
}
