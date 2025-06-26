package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterConfigurationsParameterResponse Response Object
type ListClusterConfigurationsParameterResponse struct {

	// **参数解释**： 集群使用的参数配置信息。 **取值范围**： 不涉及。
	Configurations *[]ConfigurationParameter `json:"configurations,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListClusterConfigurationsParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterConfigurationsParameterResponse struct{}"
	}

	return strings.Join([]string{"ListClusterConfigurationsParameterResponse", string(data)}, " ")
}
