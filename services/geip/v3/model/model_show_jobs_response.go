package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobsResponse Response Object
type ShowJobsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	Job *ShowJob `json:"job,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobsResponse struct{}"
	}

	return strings.Join([]string{"ShowJobsResponse", string(data)}, " ")
}
