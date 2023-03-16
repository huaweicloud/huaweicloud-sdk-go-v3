package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyResponse struct{}"
	}

	return strings.Join([]string{"DeletePolicyResponse", string(data)}, " ")
}
