package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetInstanceReadonlyStatusResponse Response Object
type SetInstanceReadonlyStatusResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetInstanceReadonlyStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceReadonlyStatusResponse struct{}"
	}

	return strings.Join([]string{"SetInstanceReadonlyStatusResponse", string(data)}, " ")
}
