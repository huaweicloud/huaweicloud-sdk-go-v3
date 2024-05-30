package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMicroserviceRouteRuleRequest Request Object
type DeleteMicroserviceRouteRuleRequest struct {

	// 该字段内容填为 \"application/json;charset=UTF-8\"。
	ContentType string `json:"Content-Type"`

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

func (o DeleteMicroserviceRouteRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMicroserviceRouteRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteMicroserviceRouteRuleRequest", string(data)}, " ")
}
