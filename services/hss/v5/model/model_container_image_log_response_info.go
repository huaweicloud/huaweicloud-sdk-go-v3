package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerImageLogResponseInfo 容器镜像使用日志信息
type ContainerImageLogResponseInfo struct {

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 事件名称
	EventName *string `json:"event_name,omitempty"`

	// 事件id
	EventId *string `json:"event_id,omitempty"`

	// 事件类型
	EventType *string `json:"event_type,omitempty"`

	// 触发事件的源ip
	SourceIp *string `json:"source_ip,omitempty"`

	// 触发事件的用户
	UserName *string `json:"user_name,omitempty"`

	// 日志产生的时间
	Time *int64 `json:"time,omitempty"`

	// 日志内容
	Content *string `json:"content,omitempty"`
}

func (o ContainerImageLogResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerImageLogResponseInfo struct{}"
	}

	return strings.Join([]string{"ContainerImageLogResponseInfo", string(data)}, " ")
}
