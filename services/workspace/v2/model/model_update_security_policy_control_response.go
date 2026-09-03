package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityPolicyControlResponse Response Object
type UpdateSecurityPolicyControlResponse struct {

	// 开启安全策略管控的资源数量。
	EnabledCount *int32 `json:"enabled_count,omitempty"`

	// 关闭安全策略管控的资源数量。
	DisabledCount *int32 `json:"disabled_count,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSecurityPolicyControlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolicyControlResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolicyControlResponse", string(data)}, " ")
}
