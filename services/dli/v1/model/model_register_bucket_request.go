package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterBucketRequest Request Object
type RegisterBucketRequest struct {
	Body *ObsBuckets `json:"body,omitempty"`
}

func (o RegisterBucketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterBucketRequest struct{}"
	}

	return strings.Join([]string{"RegisterBucketRequest", string(data)}, " ")
}
