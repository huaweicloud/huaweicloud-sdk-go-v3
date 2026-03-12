package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDryRunPolicyResponse Response Object
type DeleteDryRunPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDryRunPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDryRunPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteDryRunPolicyResponse", string(data)}, " ")
}
