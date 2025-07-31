package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPolicyResponse Response Object
type AddPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPolicyResponse struct{}"
	}

	return strings.Join([]string{"AddPolicyResponse", string(data)}, " ")
}
