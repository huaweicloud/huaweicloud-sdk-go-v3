package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseResumeDataSynchronizationRequest Request Object
type PauseResumeDataSynchronizationRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ActionBody `json:"body,omitempty"`
}

func (o PauseResumeDataSynchronizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseResumeDataSynchronizationRequest struct{}"
	}

	return strings.Join([]string{"PauseResumeDataSynchronizationRequest", string(data)}, " ")
}
