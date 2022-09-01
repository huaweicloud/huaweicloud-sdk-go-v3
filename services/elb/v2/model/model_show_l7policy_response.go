package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowL7policyResponse struct {
	L7policy       *L7policyResp `json:"l7policy,omitempty" xml:"l7policy"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowL7policyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7policyResponse struct{}"
	}

	return strings.Join([]string{"ShowL7policyResponse", string(data)}, " ")
}
