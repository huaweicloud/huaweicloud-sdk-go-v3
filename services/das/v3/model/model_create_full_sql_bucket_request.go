package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFullSqlBucketRequest Request Object
type CreateFullSqlBucketRequest struct {
	Body *CreateFullSqlBucketRequestBody `json:"body,omitempty"`
}

func (o CreateFullSqlBucketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFullSqlBucketRequest struct{}"
	}

	return strings.Join([]string{"CreateFullSqlBucketRequest", string(data)}, " ")
}
