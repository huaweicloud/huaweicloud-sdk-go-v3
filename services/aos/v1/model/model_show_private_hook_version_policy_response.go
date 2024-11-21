package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateHookVersionPolicyResponse Response Object
type ShowPrivateHookVersionPolicyResponse struct {
	Location       *string `json:"Location,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPrivateHookVersionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateHookVersionPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateHookVersionPolicyResponse", string(data)}, " ")
}
