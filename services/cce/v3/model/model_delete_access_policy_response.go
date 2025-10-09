package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccessPolicyResponse Response Object
type DeleteAccessPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAccessPolicyResponse", string(data)}, " ")
}
