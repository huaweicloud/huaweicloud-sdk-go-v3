package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrchestrationsRequest Request Object
type ListOrchestrationsRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 编排规则名称。
	OrchestrationName *string `json:"orchestration_name,omitempty"`

	// 指定需要精确匹配查找的参数名称，多个参数需要支持精确匹配时参数之间使用“,”隔开。当前仅支持orchestration_name。
	PreciseSearch *string `json:"precise_search,omitempty"`

	// 编排规则编号。  支持指定多个编号作为查询条件，多个参数之间使用“,”隔开，支持的查询参数个数与api允许绑定的参数规则上限保持一致，具体请参考产品介绍的“配额说明”章节。
	OrchestrationId *string `json:"orchestration_id,omitempty"`
}

func (o ListOrchestrationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrchestrationsRequest struct{}"
	}

	return strings.Join([]string{"ListOrchestrationsRequest", string(data)}, " ")
}
