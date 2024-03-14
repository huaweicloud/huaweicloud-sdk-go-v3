package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DelegatedAdministratorReqBody 委托管理员相关操作的请求体。
type DelegatedAdministratorReqBody struct {

	// 服务主体名称。
	ServicePrincipal string `json:"service_principal"`

	// 账号的唯一标识符（ID）。
	AccountId string `json:"account_id"`
}

func (o DelegatedAdministratorReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelegatedAdministratorReqBody struct{}"
	}

	return strings.Join([]string{"DelegatedAdministratorReqBody", string(data)}, " ")
}
