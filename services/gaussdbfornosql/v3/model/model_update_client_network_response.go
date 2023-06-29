package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClientNetworkResponse Response Object
type UpdateClientNetworkResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateClientNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientNetworkResponse struct{}"
	}

	return strings.Join([]string{"UpdateClientNetworkResponse", string(data)}, " ")
}
