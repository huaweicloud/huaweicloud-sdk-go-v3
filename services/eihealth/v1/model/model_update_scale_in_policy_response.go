package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScaleInPolicyResponse Response Object
type UpdateScaleInPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateScaleInPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScaleInPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateScaleInPolicyResponse", string(data)}, " ")
}
