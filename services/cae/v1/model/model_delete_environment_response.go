package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnvironmentResponse Response Object
type DeleteEnvironmentResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEnvironmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentResponse", string(data)}, " ")
}
