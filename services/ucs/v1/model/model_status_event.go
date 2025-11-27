package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatusEvent struct {

	// 拦截事件首次发生时间
	FirstTimestamp *string `json:"firstTimestamp,omitempty"`

	// 拦截事件资源类型
	ResourceKind *string `json:"resourceKind,omitempty"`

	// 拦截事件资源名称
	ResourceName *string `json:"resourceName,omitempty"`

	// 拦截事件资源命名空间，如果没有则为空
	ResourceNamespace *string `json:"resourceNamespace,omitempty"`

	// 拦截事件详细信息
	Message *string `json:"message,omitempty"`
}

func (o StatusEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusEvent struct{}"
	}

	return strings.Join([]string{"StatusEvent", string(data)}, " ")
}
