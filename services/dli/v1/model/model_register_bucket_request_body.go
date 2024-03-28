package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegisterBucketRequestBody struct {
	ObsBuckets []string `json:"obs_buckets"`
}

func (o RegisterBucketRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterBucketRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterBucketRequestBody", string(data)}, " ")
}
