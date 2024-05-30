package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMicroserviceRouteRuleRequest Request Object
type ListMicroserviceRouteRuleRequest struct {

	// 微服务引擎的实例ID
	XEngineId string `json:"x-engine-id"`

	// 企业项目ID
	XEnterpriseProjectID string `json:"X-Enterprise-Project-ID"`

	// 微服务名称
	ServiceName string `json:"service_name"`

	// 所属环境，不填表示<空>环境
	Environment *string `json:"environment,omitempty"`

	// 所属应用，不填默认为default应用
	AppId *string `json:"app_id,omitempty"`
}

func (o ListMicroserviceRouteRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMicroserviceRouteRuleRequest struct{}"
	}

	return strings.Join([]string{"ListMicroserviceRouteRuleRequest", string(data)}, " ")
}
