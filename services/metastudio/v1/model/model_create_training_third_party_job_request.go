package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainingThirdPartyJobRequest Request Object
type CreateTrainingThirdPartyJobRequest struct {

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	Body *CreateTrainingJobReq `json:"body,omitempty"`
}

func (o CreateTrainingThirdPartyJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingThirdPartyJobRequest struct{}"
	}

	return strings.Join([]string{"CreateTrainingThirdPartyJobRequest", string(data)}, " ")
}
