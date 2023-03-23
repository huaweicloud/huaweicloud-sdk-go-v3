package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisableOrganizationShareResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableOrganizationShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableOrganizationShareResponse struct{}"
	}

	return strings.Join([]string{"DisableOrganizationShareResponse", string(data)}, " ")
}
