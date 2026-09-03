package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleTemplateTrendRequest Request Object
type ShowSingleTemplateTrendRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 模板ID
	TemplateId string `json:"template_id"`

	// 开始时间，单位毫秒
	StartAt int64 `json:"start_at"`

	// 结束时间，单位毫秒
	EndAt int64 `json:"end_at"`

	// 聚合毫秒数
	IntervalMillis *int64 `json:"interval_millis,omitempty"`
}

func (o ShowSingleTemplateTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleTemplateTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowSingleTemplateTrendRequest", string(data)}, " ")
}
