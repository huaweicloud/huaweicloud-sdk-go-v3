package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeRunInfoInnerError struct {

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	ErrorBody *NodeRunInfoInnerErrorErrorBody `json:"error_body,omitempty"`
}

func (o NodeRunInfoInnerError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeRunInfoInnerError struct{}"
	}

	return strings.Join([]string{"NodeRunInfoInnerError", string(data)}, " ")
}
