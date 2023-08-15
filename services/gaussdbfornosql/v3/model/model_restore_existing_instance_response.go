package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreExistingInstanceResponse Response Object
type RestoreExistingInstanceResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreExistingInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreExistingInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreExistingInstanceResponse", string(data)}, " ")
}
