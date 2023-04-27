package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartScaleOutPolicyRequest struct {

	// 策略id
	Id string `json:"id"`
}

func (o StartScaleOutPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartScaleOutPolicyRequest struct{}"
	}

	return strings.Join([]string{"StartScaleOutPolicyRequest", string(data)}, " ")
}
