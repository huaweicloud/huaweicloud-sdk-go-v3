package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkloadIdentityResponse Response Object
type DeleteWorkloadIdentityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWorkloadIdentityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkloadIdentityResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkloadIdentityResponse", string(data)}, " ")
}
