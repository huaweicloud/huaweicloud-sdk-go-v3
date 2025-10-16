package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterSupportConfigurationResponse Response Object
type ShowClusterSupportConfigurationResponse struct {

	// **参数解释**： 集群支持的配置项详情
	Body           map[string][]PackageOptions `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowClusterSupportConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterSupportConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterSupportConfigurationResponse", string(data)}, " ")
}
