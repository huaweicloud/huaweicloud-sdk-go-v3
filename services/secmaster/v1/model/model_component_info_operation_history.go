package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentInfoOperationHistory struct {

	// 操作
	OperationName *string `json:"operation_name,omitempty"`

	// 时间点
	OperationTime *string `json:"operation_time,omitempty"`
}

func (o ComponentInfoOperationHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentInfoOperationHistory struct{}"
	}

	return strings.Join([]string{"ComponentInfoOperationHistory", string(data)}, " ")
}
