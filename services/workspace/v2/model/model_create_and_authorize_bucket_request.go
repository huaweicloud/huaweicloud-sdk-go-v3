package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAndAuthorizeBucketRequest Request Object
type CreateAndAuthorizeBucketRequest struct {
}

func (o CreateAndAuthorizeBucketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAndAuthorizeBucketRequest struct{}"
	}

	return strings.Join([]string{"CreateAndAuthorizeBucketRequest", string(data)}, " ")
}
