package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteOptions struct {

	// **参数解释**：序列号。
	SerialNumber *string `json:"serialNumber,omitempty"`
}

func (o ExecuteOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteOptions struct{}"
	}

	return strings.Join([]string{"ExecuteOptions", string(data)}, " ")
}
