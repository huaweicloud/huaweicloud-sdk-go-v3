package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAiGuardPolicyResponse Response Object
type DeleteAiGuardPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAiGuardPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAiGuardPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAiGuardPolicyResponse", string(data)}, " ")
}
