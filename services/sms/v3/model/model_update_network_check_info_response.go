package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNetworkCheckInfoResponse Response Object
type UpdateNetworkCheckInfoResponse struct {

	// 任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNetworkCheckInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkCheckInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateNetworkCheckInfoResponse", string(data)}, " ")
}
