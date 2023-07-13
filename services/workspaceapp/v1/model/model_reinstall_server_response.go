package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReinstallServerResponse Response Object
type ReinstallServerResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ReinstallServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerResponse struct{}"
	}

	return strings.Join([]string{"ReinstallServerResponse", string(data)}, " ")
}
