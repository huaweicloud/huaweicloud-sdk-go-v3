package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteOrganizationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOrganizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationResponse struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationResponse", string(data)}, " ")
}
