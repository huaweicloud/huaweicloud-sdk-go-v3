package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScaleOutPolicyResponse Response Object
type CreateScaleOutPolicyResponse struct {

	// 扩容策略id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScaleOutPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScaleOutPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateScaleOutPolicyResponse", string(data)}, " ")
}
