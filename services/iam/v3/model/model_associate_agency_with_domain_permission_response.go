package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateAgencyWithDomainPermissionResponse Response Object
type AssociateAgencyWithDomainPermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateAgencyWithDomainPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAgencyWithDomainPermissionResponse struct{}"
	}

	return strings.Join([]string{"AssociateAgencyWithDomainPermissionResponse", string(data)}, " ")
}
