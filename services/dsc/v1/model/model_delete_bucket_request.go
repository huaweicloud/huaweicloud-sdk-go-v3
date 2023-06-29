package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBucketRequest Request Object
type DeleteBucketRequest struct {

	// 桶ID
	BucketId string `json:"bucket_id"`
}

func (o DeleteBucketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBucketRequest struct{}"
	}

	return strings.Join([]string{"DeleteBucketRequest", string(data)}, " ")
}
