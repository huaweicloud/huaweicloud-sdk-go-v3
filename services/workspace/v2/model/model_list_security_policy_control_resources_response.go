package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityPolicyControlResourcesResponse Response Object
type ListSecurityPolicyControlResourcesResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 安全策略管控资源列表项。
	Items *[]SecurityPolicyControlResourceItemVo `json:"items,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSecurityPolicyControlResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPolicyControlResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityPolicyControlResourcesResponse", string(data)}, " ")
}
