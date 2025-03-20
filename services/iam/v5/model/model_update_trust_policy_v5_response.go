package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTrustPolicyV5Response Response Object
type UpdateTrustPolicyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTrustPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrustPolicyV5Response struct{}"
	}

	return strings.Join([]string{"UpdateTrustPolicyV5Response", string(data)}, " ")
}
