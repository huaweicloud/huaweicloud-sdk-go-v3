package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateL7policiesRequestBody struct {
	L7policy *UpdateL7policyReq `json:"l7policy" xml:"l7policy"`
}

func (o UpdateL7policiesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7policiesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateL7policiesRequestBody", string(data)}, " ")
}
