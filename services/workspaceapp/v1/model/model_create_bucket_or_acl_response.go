package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBucketOrAclResponse Response Object
type CreateBucketOrAclResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateBucketOrAclResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBucketOrAclResponse struct{}"
	}

	return strings.Join([]string{"CreateBucketOrAclResponse", string(data)}, " ")
}
