package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainingAdvanceJobRequest Request Object
type CreateTrainingAdvanceJobRequest struct {

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	Body *CreateTrainingJobReq `json:"body,omitempty"`
}

func (o CreateTrainingAdvanceJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingAdvanceJobRequest struct{}"
	}

	return strings.Join([]string{"CreateTrainingAdvanceJobRequest", string(data)}, " ")
}
