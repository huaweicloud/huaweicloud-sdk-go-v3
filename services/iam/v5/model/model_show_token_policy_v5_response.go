package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTokenPolicyV5Response Response Object
type ShowTokenPolicyV5Response struct {
	TokenPolicy    *TokenPolicy `json:"token_policy,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowTokenPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTokenPolicyV5Response struct{}"
	}

	return strings.Join([]string{"ShowTokenPolicyV5Response", string(data)}, " ")
}
