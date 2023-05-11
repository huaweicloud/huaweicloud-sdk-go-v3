package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEditingJobsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEditingJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEditingJobsResponse struct{}"
	}

	return strings.Join([]string{"DeleteEditingJobsResponse", string(data)}, " ")
}
