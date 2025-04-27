package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyAutoExpandPolicyResponse Response Object
type ModifyAutoExpandPolicyResponse struct {

	// 修改结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyAutoExpandPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAutoExpandPolicyResponse struct{}"
	}

	return strings.Join([]string{"ModifyAutoExpandPolicyResponse", string(data)}, " ")
}
