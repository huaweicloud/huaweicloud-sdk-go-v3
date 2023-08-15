package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsObjectRequest Request Object
type ListObsObjectRequest struct {

	// obs桶名称
	BucketName string `json:"bucket_name"`

	// 查询起始object名称
	Marker *string `json:"marker,omitempty"`

	// 单次查询条数
	Limit int32 `json:"limit"`

	// 搜索前缀
	Prefix *string `json:"prefix,omitempty"`
}

func (o ListObsObjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsObjectRequest struct{}"
	}

	return strings.Join([]string{"ListObsObjectRequest", string(data)}, " ")
}
