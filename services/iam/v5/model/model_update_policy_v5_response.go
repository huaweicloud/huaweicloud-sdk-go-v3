package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyV5Response Response Object
type UpdatePolicyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyV5Response struct{}"
	}

	return strings.Join([]string{"UpdatePolicyV5Response", string(data)}, " ")
}
