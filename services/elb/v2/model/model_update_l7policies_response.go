package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateL7policiesResponse struct {
	L7policy       *L7policyResp `json:"l7policy,omitempty" xml:"l7policy"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateL7policiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7policiesResponse struct{}"
	}

	return strings.Join([]string{"UpdateL7policiesResponse", string(data)}, " ")
}
