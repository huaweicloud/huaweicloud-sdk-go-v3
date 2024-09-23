package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterDelegatedAdministratorRequest Request Object
type RegisterDelegatedAdministratorRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	Body *DelegatedAdministratorReqBody `json:"body,omitempty"`
}

func (o RegisterDelegatedAdministratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDelegatedAdministratorRequest struct{}"
	}

	return strings.Join([]string{"RegisterDelegatedAdministratorRequest", string(data)}, " ")
}
