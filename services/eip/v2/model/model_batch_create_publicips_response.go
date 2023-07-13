package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreatePublicipsResponse Response Object
type BatchCreatePublicipsResponse struct {

	// job_id，需要访问调用netAPI组件访问job执行情况。netAPI：/v1/{tenant_id}/jobs/{job_id}
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreatePublicipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicipsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicipsResponse", string(data)}, " ")
}
