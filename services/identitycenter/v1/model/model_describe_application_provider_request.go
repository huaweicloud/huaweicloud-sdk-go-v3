package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeApplicationProviderRequest Request Object
type DescribeApplicationProviderRequest struct {

	// 应用程序提供方Id
	ApplicationProviderId string `json:"application_provider_id"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o DescribeApplicationProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeApplicationProviderRequest struct{}"
	}

	return strings.Join([]string{"DescribeApplicationProviderRequest", string(data)}, " ")
}
