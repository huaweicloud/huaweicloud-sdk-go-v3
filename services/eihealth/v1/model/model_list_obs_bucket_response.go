package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsBucketResponse Response Object
type ListObsBucketResponse struct {

	// 桶个数
	Count *int32 `json:"count,omitempty"`

	// 桶列表
	Buckets        *[]BucketDto `json:"buckets,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListObsBucketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketResponse struct{}"
	}

	return strings.Join([]string{"ListObsBucketResponse", string(data)}, " ")
}
