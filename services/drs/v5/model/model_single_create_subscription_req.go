package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SingleCreateSubscriptionReq 创建单个订阅任务的请求体
type SingleCreateSubscriptionReq struct {

	// 任务名称 约束：任务名称在4位到50位之间，不区分大小写，可以包含字母、数字、中划线或下划线，不能包括其他特殊字符。 - 最小长度：4 - 最大长度：50
	Name string `json:"name"`

	// 任务描述
	Description *string `json:"description,omitempty"`

	// 实例类型，仅支持rds
	InstanceType *string `json:"instance_type,omitempty"`

	// 企业项目id
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 标签
	Tags *[]ResourceTag `json:"tags,omitempty"`

	SourceEndpointInfo *SubscriptionSourceEndpointInfo `json:"source_endpoint_info"`

	// 是否创建委托，取值： - true：创建 - false：不创建 默认为false
	IsGrantNewAgency *bool `json:"is_grant_new_agency,omitempty"`
}

func (o SingleCreateSubscriptionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleCreateSubscriptionReq struct{}"
	}

	return strings.Join([]string{"SingleCreateSubscriptionReq", string(data)}, " ")
}
