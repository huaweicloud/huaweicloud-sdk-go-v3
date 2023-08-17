package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsBucketsResponse Response Object
type ListObsBucketsResponse struct {

	// 桶的总数
	Count *int32 `json:"count,omitempty"`

	// 桶信息
	Buckets *[]ObsBucket `json:"buckets,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListObsBucketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketsResponse struct{}"
	}

	return strings.Join([]string{"ListObsBucketsResponse", string(data)}, " ")
}
