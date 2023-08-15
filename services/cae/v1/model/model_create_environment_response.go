package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnvironmentResponse Response Object
type CreateEnvironmentResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEnvironmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentResponse", string(data)}, " ")
}
