package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDecoyPortPolicyResponse Response Object
type CreateDecoyPortPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDecoyPortPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDecoyPortPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateDecoyPortPolicyResponse", string(data)}, " ")
}
