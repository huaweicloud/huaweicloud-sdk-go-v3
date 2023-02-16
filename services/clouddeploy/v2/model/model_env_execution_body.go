package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用启动参数
type EnvExecutionBody struct {

	// 部署应用时传递的参数
	Params *[]DynamicConfigInfo `json:"params,omitempty"`

	// 应用的部署id，可通过record_id回滚至之前的部署状态。选中应用历史部署记录，在URL中获取
	RecordId *string `json:"record_id,omitempty"`

	// 限制触发来源,0不限制任何部署请求来源,1时只允许通过流水线触发部署
	TriggerSource *string `json:"trigger_source,omitempty"`
}

func (o EnvExecutionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvExecutionBody struct{}"
	}

	return strings.Join([]string{"EnvExecutionBody", string(data)}, " ")
}
