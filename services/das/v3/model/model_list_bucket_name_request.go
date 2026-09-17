package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBucketNameRequest Request Object
type ListBucketNameRequest struct {
}

func (o ListBucketNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBucketNameRequest struct{}"
	}

	return strings.Join([]string{"ListBucketNameRequest", string(data)}, " ")
}
