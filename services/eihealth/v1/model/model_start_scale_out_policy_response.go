package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartScaleOutPolicyResponse Response Object
type StartScaleOutPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartScaleOutPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartScaleOutPolicyResponse struct{}"
	}

	return strings.Join([]string{"StartScaleOutPolicyResponse", string(data)}, " ")
}
