package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetClusterSupportConfigurationResponse Response Object
type GetClusterSupportConfigurationResponse struct {

	// **参数解释**： 集群支持的配置项详情
	Body           map[string][]PackageOptions `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o GetClusterSupportConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetClusterSupportConfigurationResponse struct{}"
	}

	return strings.Join([]string{"GetClusterSupportConfigurationResponse", string(data)}, " ")
}
