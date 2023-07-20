package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopFlinkJobsResponse Response Object
type StopFlinkJobsResponse struct {
	Body           *[]CommonResp `json:"body,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o StopFlinkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopFlinkJobsResponse struct{}"
	}

	return strings.Join([]string{"StopFlinkJobsResponse", string(data)}, " ")
}
