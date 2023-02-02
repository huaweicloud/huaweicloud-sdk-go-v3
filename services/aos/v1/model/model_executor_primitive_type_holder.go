package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutorPrimitiveTypeHolder struct {
}

func (o ExecutorPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutorPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ExecutorPrimitiveTypeHolder", string(data)}, " ")
}
