package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainsNewRequest Request Object
type ShowDomainsNewRequest struct {

	// 认证类型。 - OPEN_API：UOS域控。
	AuthType *string `json:"auth_type,omitempty"`
}

func (o ShowDomainsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainsNewRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainsNewRequest", string(data)}, " ")
}
