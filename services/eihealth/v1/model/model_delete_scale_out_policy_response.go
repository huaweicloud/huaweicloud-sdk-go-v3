package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScaleOutPolicyResponse Response Object
type DeleteScaleOutPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScaleOutPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScaleOutPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteScaleOutPolicyResponse", string(data)}, " ")
}
