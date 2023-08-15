package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTtsaResponse Response Object
type CreateTtsaResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTtsaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtsaResponse struct{}"
	}

	return strings.Join([]string{"CreateTtsaResponse", string(data)}, " ")
}
