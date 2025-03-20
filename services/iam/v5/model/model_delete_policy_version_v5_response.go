package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyVersionV5Response Response Object
type DeletePolicyVersionV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePolicyVersionV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyVersionV5Response struct{}"
	}

	return strings.Join([]string{"DeletePolicyVersionV5Response", string(data)}, " ")
}
