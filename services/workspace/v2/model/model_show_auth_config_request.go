package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuthConfigRequest Request Object
type ShowAuthConfigRequest struct {

	// 认证类型
	AuthType *string `json:"auth_type,omitempty"`
}

func (o ShowAuthConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAuthConfigRequest", string(data)}, " ")
}
