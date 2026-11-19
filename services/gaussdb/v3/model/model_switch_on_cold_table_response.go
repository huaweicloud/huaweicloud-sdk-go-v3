package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchOnColdTableResponse Response Object
type SwitchOnColdTableResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchOnColdTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchOnColdTableResponse struct{}"
	}

	return strings.Join([]string{"SwitchOnColdTableResponse", string(data)}, " ")
}
