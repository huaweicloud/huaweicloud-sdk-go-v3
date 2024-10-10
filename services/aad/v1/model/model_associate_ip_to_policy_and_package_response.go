package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIpToPolicyAndPackageResponse Response Object
type AssociateIpToPolicyAndPackageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateIpToPolicyAndPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIpToPolicyAndPackageResponse struct{}"
	}

	return strings.Join([]string{"AssociateIpToPolicyAndPackageResponse", string(data)}, " ")
}
