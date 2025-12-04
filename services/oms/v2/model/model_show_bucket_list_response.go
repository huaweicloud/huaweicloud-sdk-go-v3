package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBucketListResponse Response Object
type ShowBucketListResponse struct {

	// 桶名列表
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowBucketListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBucketListResponse struct{}"
	}

	return strings.Join([]string{"ShowBucketListResponse", string(data)}, " ")
}
