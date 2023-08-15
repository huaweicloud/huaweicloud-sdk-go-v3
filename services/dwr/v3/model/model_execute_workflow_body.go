package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteWorkflowBody struct {

	// 桶名
	Bucket string `json:"bucket"`

	// 对象名
	Object string `json:"object"`

	Inputs *Input `json:"inputs,omitempty"`
}

func (o ExecuteWorkflowBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteWorkflowBody struct{}"
	}

	return strings.Join([]string{"ExecuteWorkflowBody", string(data)}, " ")
}
