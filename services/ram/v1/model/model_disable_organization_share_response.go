package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableOrganizationShareResponse Response Object
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
