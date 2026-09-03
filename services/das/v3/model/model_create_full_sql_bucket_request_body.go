package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFullSqlBucketRequestBody 创建全量SQL桶请求体
type CreateFullSqlBucketRequestBody struct {

	// OBS桶名
	BucketName string `json:"bucket_name"`
}

func (o CreateFullSqlBucketRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFullSqlBucketRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFullSqlBucketRequestBody", string(data)}, " ")
}
