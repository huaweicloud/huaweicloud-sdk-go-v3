package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreOldInstanceResponse Response Object
type RestoreOldInstanceResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreOldInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreOldInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreOldInstanceResponse", string(data)}, " ")
}
