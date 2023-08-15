package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScaleOutPolicyResponse Response Object
type UpdateScaleOutPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateScaleOutPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScaleOutPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateScaleOutPolicyResponse", string(data)}, " ")
}
