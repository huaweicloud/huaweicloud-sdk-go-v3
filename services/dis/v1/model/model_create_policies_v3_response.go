package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePoliciesV3Response struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePoliciesV3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoliciesV3Response struct{}"
	}

	return strings.Join([]string{"CreatePoliciesV3Response", string(data)}, " ")
}
