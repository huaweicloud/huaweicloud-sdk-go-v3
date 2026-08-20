package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessInstanceResponseResultStatus 评审单状态
type ProcessInstanceResponseResultStatus struct {

	// 状态码
	Code *string `json:"code,omitempty"`

	// 状态
	Name *string `json:"name,omitempty"`
}

func (o ProcessInstanceResponseResultStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInstanceResponseResultStatus struct{}"
	}

	return strings.Join([]string{"ProcessInstanceResponseResultStatus", string(data)}, " ")
}
