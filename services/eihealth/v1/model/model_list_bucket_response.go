package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBucketResponse Response Object
type ListBucketResponse struct {

	// 桶个数
	Count *int32 `json:"count,omitempty"`

	// 桶列表
	Buckets        *[]ProjectBucketRsp `json:"buckets,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListBucketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBucketResponse struct{}"
	}

	return strings.Join([]string{"ListBucketResponse", string(data)}, " ")
}
