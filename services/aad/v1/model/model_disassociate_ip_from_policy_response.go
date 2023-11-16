package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateIpFromPolicyResponse Response Object
type DisassociateIpFromPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateIpFromPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateIpFromPolicyResponse struct{}"
	}

	return strings.Join([]string{"DisassociateIpFromPolicyResponse", string(data)}, " ")
}
