package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BucketDto OBS桶
type BucketDto struct {

	// 桶名称
	Name *string `json:"name,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	Type *BucketType `json:"type,omitempty"`
}

func (o BucketDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BucketDto struct{}"
	}

	return strings.Join([]string{"BucketDto", string(data)}, " ")
}
