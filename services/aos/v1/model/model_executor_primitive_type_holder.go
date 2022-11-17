package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutorPrimitiveTypeHolder struct {

	// 执行操作者的名字，将用做未来的审计工作
	Executor *string `json:"executor,omitempty"`
}

func (o ExecutorPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutorPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ExecutorPrimitiveTypeHolder", string(data)}, " ")
}
