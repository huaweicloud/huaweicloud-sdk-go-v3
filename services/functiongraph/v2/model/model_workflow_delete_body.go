package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量删除函数流的Body体
type WorkflowDeleteBody struct {

	// 函数流URN列表
	WorkflowUrns []string `json:"workflow_urns"`
}

func (o WorkflowDeleteBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowDeleteBody struct{}"
	}

	return strings.Join([]string{"WorkflowDeleteBody", string(data)}, " ")
}
