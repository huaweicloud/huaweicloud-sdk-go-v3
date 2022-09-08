package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部署任务启动参数
type EnvExecutionBody struct {

	// 部署任务执行时传递的参数
	Params []DynamicConfigInfo `json:"params"`

	// 部署任务的执行id，可通过record_id回滚至之前的部署状态。选中部署历史执行记录，在URL中获取。
	RecordId *string `json:"record_id,omitempty"`
}

func (o EnvExecutionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvExecutionBody struct{}"
	}

	return strings.Join([]string{"EnvExecutionBody", string(data)}, " ")
}
