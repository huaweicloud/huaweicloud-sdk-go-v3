package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWorkflowBody struct {

	// 工作流中每个算子待修改参数与待更新的值。
	Parameters map[string]string `json:"parameters,omitempty"`
}

func (o UpdateWorkflowBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowBody struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowBody", string(data)}, " ")
}
