package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyV5Response Response Object
type DeletePolicyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyV5Response struct{}"
	}

	return strings.Join([]string{"DeletePolicyV5Response", string(data)}, " ")
}
