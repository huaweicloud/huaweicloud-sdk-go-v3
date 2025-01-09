package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAndAuthorizeBucketResponse Response Object
type CreateAndAuthorizeBucketResponse struct {

	// 桶名称。
	BucketName     *string `json:"bucket_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAndAuthorizeBucketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAndAuthorizeBucketResponse struct{}"
	}

	return strings.Join([]string{"CreateAndAuthorizeBucketResponse", string(data)}, " ")
}
