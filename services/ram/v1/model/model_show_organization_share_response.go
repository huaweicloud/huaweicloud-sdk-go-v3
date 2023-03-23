package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowOrganizationShareResponse struct {

	// 如果为\"true\"，则表示启用与组织的共享。默认为\"false\"。
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowOrganizationShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationShareResponse struct{}"
	}

	return strings.Join([]string{"ShowOrganizationShareResponse", string(data)}, " ")
}
