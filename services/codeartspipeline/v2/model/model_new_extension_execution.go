package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NewExtensionExecution 执行信息
type NewExtensionExecution struct {

	// 入口
	Target *string `json:"target,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// sha256
	Sha256 *string `json:"sha256,omitempty"`
}

func (o NewExtensionExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewExtensionExecution struct{}"
	}

	return strings.Join([]string{"NewExtensionExecution", string(data)}, " ")
}
