package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsBucketsResponse Response Object
type ListObsBucketsResponse struct {

	// obs桶列表
	Buckets *[]BucketDetail `json:"buckets,omitempty"`

	// obs桶总数
	Total *int32 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListObsBucketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketsResponse struct{}"
	}

	return strings.Join([]string{"ListObsBucketsResponse", string(data)}, " ")
}
