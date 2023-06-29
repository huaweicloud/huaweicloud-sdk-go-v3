package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBucketsResponse Response Object
type ListBucketsResponse struct {

	// OBS桶列表
	Buckets *[]Bucket `json:"buckets,omitempty"`

	// OBS桶总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBucketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBucketsResponse struct{}"
	}

	return strings.Join([]string{"ListBucketsResponse", string(data)}, " ")
}
