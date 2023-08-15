package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDomainPermissionForAgencyResponse Response Object
type CheckDomainPermissionForAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckDomainPermissionForAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDomainPermissionForAgencyResponse struct{}"
	}

	return strings.Join([]string{"CheckDomainPermissionForAgencyResponse", string(data)}, " ")
}
