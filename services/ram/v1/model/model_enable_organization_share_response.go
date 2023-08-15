package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableOrganizationShareResponse Response Object
type EnableOrganizationShareResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableOrganizationShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableOrganizationShareResponse struct{}"
	}

	return strings.Join([]string{"EnableOrganizationShareResponse", string(data)}, " ")
}
