package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterSecurityConfigurationsResponse Response Object
type ListClusterSecurityConfigurationsResponse struct {

	// **参数解释**： 参数列表。 **取值范围**： 不涉及。
	Configurations *[]SecurityConfigurationParameter `json:"configurations,omitempty"`

	// **参数解释**： 总条数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClusterSecurityConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterSecurityConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterSecurityConfigurationsResponse", string(data)}, " ")
}
