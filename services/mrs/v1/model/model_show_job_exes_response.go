package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobExesResponse struct {
	JobExecution   *JobExeResult `json:"job_execution,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowJobExesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobExesResponse struct{}"
	}

	return strings.Join([]string{"ShowJobExesResponse", string(data)}, " ")
}
