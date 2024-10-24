package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationShareRequest Request Object
type ShowOrganizationShareRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o ShowOrganizationShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationShareRequest struct{}"
	}

	return strings.Join([]string{"ShowOrganizationShareRequest", string(data)}, " ")
}
