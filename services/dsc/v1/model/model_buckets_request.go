package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BucketsRequest struct {

	// OBS桶列表
	Buckets *[]BucketBean `json:"buckets,omitempty"`
}

func (o BucketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BucketsRequest struct{}"
	}

	return strings.Join([]string{"BucketsRequest", string(data)}, " ")
}
