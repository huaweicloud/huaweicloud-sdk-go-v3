package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateOrganizationResponse struct {
	Organization   *OrganizationDto `json:"organization,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateOrganizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationResponse struct{}"
	}

	return strings.Join([]string{"CreateOrganizationResponse", string(data)}, " ")
}
