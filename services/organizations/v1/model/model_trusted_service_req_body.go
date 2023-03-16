package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 可信服务相关操作的请求体。
type TrustedServiceReqBody struct {

	// 服务主体名称。
	ServicePrincipal string `json:"service_principal"`
}

func (o TrustedServiceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrustedServiceReqBody struct{}"
	}

	return strings.Join([]string{"TrustedServiceReqBody", string(data)}, " ")
}
