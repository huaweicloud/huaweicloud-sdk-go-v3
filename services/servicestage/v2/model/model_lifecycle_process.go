package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LifecycleProcess struct {

	// 取值为command或者http。command为执行命令行，http为发送http请求。
	Type *string `json:"type,omitempty"`

	Parameters *LifecycleProcessParameter `json:"parameters,omitempty"`
}

func (o LifecycleProcess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifecycleProcess struct{}"
	}

	return strings.Join([]string{"LifecycleProcess", string(data)}, " ")
}
