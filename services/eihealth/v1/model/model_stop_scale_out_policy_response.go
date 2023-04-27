package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopScaleOutPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopScaleOutPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopScaleOutPolicyResponse struct{}"
	}

	return strings.Join([]string{"StopScaleOutPolicyResponse", string(data)}, " ")
}
