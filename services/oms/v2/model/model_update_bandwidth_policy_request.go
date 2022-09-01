package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateBandwidthPolicyRequest struct {

	// 任务ID。
	TaskId int64 `json:"task_id" xml:"task_id"`

	Body *UpdateBandwidthPolicyReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateBandwidthPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthPolicyRequest", string(data)}, " ")
}
