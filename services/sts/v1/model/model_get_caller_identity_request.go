package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetCallerIdentityRequest Request Object
type GetCallerIdentityRequest struct {

	// 通过临时访问密钥调用接口时，需要提供“X-Security-Token”Http头，取值为临时访问密钥的security_token字段。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o GetCallerIdentityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetCallerIdentityRequest struct{}"
	}

	return strings.Join([]string{"GetCallerIdentityRequest", string(data)}, " ")
}
