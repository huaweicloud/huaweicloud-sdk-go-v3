package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceSignPolicyResponse Response Object
type DeleteInstanceSignPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceSignPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceSignPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceSignPolicyResponse", string(data)}, " ")
}
