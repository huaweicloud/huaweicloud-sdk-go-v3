package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTokenPolicyV5Response Response Object
type UpdateTokenPolicyV5Response struct {
	TokenPolicy    *TokenPolicy `json:"token_policy,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateTokenPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTokenPolicyV5Response struct{}"
	}

	return strings.Join([]string{"UpdateTokenPolicyV5Response", string(data)}, " ")
}
