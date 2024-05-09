package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteFlinkJobsResponse Response Object
type BatchDeleteFlinkJobsResponse struct {
	Body           *[]FlinkSuccessResponse `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchDeleteFlinkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFlinkJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteFlinkJobsResponse", string(data)}, " ")
}
