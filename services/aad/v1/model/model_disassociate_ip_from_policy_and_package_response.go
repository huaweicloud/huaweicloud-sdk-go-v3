package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateIpFromPolicyAndPackageResponse Response Object
type DisassociateIpFromPolicyAndPackageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateIpFromPolicyAndPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateIpFromPolicyAndPackageResponse struct{}"
	}

	return strings.Join([]string{"DisassociateIpFromPolicyAndPackageResponse", string(data)}, " ")
}
