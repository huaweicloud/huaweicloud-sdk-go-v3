package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGovernancePolicysRequest Request Object
type ListGovernancePolicysRequest struct {

	// 该字段内容填为 \"application/json;charset=UTF-8\"。
	ContentType string `json:"Content-Type"`

	// 微服务引擎的实例ID
	XEngineId string `json:"x-engine-id"`

	// 企业项目ID
	XEnterpriseProjectID string `json:"X-Enterprise-Project-ID"`

	// 所属环境，填写all时表示查询所有环境。
	Environment string `json:"environment"`

	// 所属应用
	App *string `json:"app,omitempty"`
}

func (o ListGovernancePolicysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGovernancePolicysRequest struct{}"
	}

	return strings.Join([]string{"ListGovernancePolicysRequest", string(data)}, " ")
}
