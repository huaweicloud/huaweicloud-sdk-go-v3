package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindEipResponse Response Object
type UnbindEipResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnbindEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindEipResponse struct{}"
	}

	return strings.Join([]string{"UnbindEipResponse", string(data)}, " ")
}
