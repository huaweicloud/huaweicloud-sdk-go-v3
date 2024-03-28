package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterBucketRequest Request Object
type RegisterBucketRequest struct {
	Body *RegisterBucketRequestBody `json:"body,omitempty"`
}

func (o RegisterBucketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterBucketRequest struct{}"
	}

	return strings.Join([]string{"RegisterBucketRequest", string(data)}, " ")
}
