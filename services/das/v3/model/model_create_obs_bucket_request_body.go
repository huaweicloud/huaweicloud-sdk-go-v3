package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObsBucketRequestBody Create Obs Bucket请求体
type CreateObsBucketRequestBody struct {

	// OBS桶名
	BucketName string `json:"bucket_name"`
}

func (o CreateObsBucketRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObsBucketRequestBody struct{}"
	}

	return strings.Join([]string{"CreateObsBucketRequestBody", string(data)}, " ")
}
