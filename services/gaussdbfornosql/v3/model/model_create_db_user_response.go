package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbUserResponse Response Object
type CreateDbUserResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDbUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbUserResponse struct{}"
	}

	return strings.Join([]string{"CreateDbUserResponse", string(data)}, " ")
}
