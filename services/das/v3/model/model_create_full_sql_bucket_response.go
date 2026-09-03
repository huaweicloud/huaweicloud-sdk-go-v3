package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFullSqlBucketResponse Response Object
type CreateFullSqlBucketResponse struct {

	// 是否创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateFullSqlBucketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFullSqlBucketResponse struct{}"
	}

	return strings.Join([]string{"CreateFullSqlBucketResponse", string(data)}, " ")
}
