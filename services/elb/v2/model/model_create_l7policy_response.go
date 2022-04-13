package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateL7policyResponse struct {
	L7policy       *L7policyResp `json:"l7policy,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateL7policyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7policyResponse struct{}"
	}

	return strings.Join([]string{"CreateL7policyResponse", string(data)}, " ")
}
