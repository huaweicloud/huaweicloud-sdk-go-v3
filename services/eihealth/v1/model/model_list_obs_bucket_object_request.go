package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListObsBucketObjectRequest struct {

	// 桶名称
	BucketName string `json:"bucket_name"`

	// 限制量，单次查询总量[1, 1000]，默认100
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 子路径
	Path *string `json:"path,omitempty"`
}

func (o ListObsBucketObjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketObjectRequest struct{}"
	}

	return strings.Join([]string{"ListObsBucketObjectRequest", string(data)}, " ")
}
