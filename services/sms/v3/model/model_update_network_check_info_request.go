package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNetworkCheckInfoRequest Request Object
type UpdateNetworkCheckInfoRequest struct {

	// 任务id
	TaskId string `json:"task_id"`

	Body *NetworkCheckInfoRequestBody `json:"body,omitempty"`
}

func (o UpdateNetworkCheckInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkCheckInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateNetworkCheckInfoRequest", string(data)}, " ")
}
