package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationResponse Response Object
type ShowOrganizationResponse struct {
	Organization   *OrganizationDto `json:"organization,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowOrganizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationResponse struct{}"
	}

	return strings.Join([]string{"ShowOrganizationResponse", string(data)}, " ")
}
