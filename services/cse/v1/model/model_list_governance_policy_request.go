package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGovernancePolicyRequest Request Object
type ListGovernancePolicyRequest struct {

	// 该字段内容填为 \"application/json;charset=UTF-8\"。
	ContentType string `json:"Content-Type"`

	// 微服务引擎的实例ID
	XEngineId string `json:"x-engine-id"`

	// 企业项目ID
	XEnterpriseProjectID string `json:"X-Enterprise-Project-ID"`

	// 所属环境
	XEnvironment *string `json:"x-environment,omitempty"`

	// 治理策略类型
	Kind string `json:"kind"`
}

func (o ListGovernancePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGovernancePolicyRequest struct{}"
	}

	return strings.Join([]string{"ListGovernancePolicyRequest", string(data)}, " ")
}
