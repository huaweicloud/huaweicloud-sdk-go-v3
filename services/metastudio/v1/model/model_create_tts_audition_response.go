package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTtsAuditionResponse Response Object
type CreateTtsAuditionResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTtsAuditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtsAuditionResponse struct{}"
	}

	return strings.Join([]string{"CreateTtsAuditionResponse", string(data)}, " ")
}
