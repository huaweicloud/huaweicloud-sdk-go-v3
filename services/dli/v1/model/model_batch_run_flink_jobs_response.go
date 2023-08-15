package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRunFlinkJobsResponse Response Object
type BatchRunFlinkJobsResponse struct {
	Body           *[]CommonResp `json:"body,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o BatchRunFlinkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRunFlinkJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchRunFlinkJobsResponse", string(data)}, " ")
}
