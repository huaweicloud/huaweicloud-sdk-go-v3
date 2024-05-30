package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGovernancePolicyRequest Request Object
type CreateGovernancePolicyRequest struct {

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

	Body *CreateGovPolicy `json:"body,omitempty"`
}

func (o CreateGovernancePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGovernancePolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateGovernancePolicyRequest", string(data)}, " ")
}
