package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBucketOrAclRequest Request Object
type CreateBucketOrAclRequest struct {
}

func (o CreateBucketOrAclRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBucketOrAclRequest struct{}"
	}

	return strings.Join([]string{"CreateBucketOrAclRequest", string(data)}, " ")
}
