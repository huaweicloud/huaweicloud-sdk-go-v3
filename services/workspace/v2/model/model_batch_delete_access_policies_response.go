package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteAccessPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteAccessPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAccessPoliciesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAccessPoliciesResponse", string(data)}, " ")
}
