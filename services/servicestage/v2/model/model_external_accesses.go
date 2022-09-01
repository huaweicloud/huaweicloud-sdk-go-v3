package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 访问方式。
type ExternalAccesses struct {

	// ID。
	Id *string `json:"id,omitempty" xml:"id"`

	Protocol *ExternalAccessProtocol `json:"protocol" xml:"protocol"`

	// 访问地址。
	Address string `json:"address" xml:"address"`

	// 应用组件进程监听端口
	ForwardPort int32 `json:"forward_port" xml:"forward_port"`

	Type *ExternalAccessType `json:"type,omitempty" xml:"type"`

	Status *ExternalAccessStatus `json:"status,omitempty" xml:"status"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 修改时间。
	UpdateTime *int64 `json:"update_time,omitempty" xml:"update_time"`
}

func (o ExternalAccesses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalAccesses struct{}"
	}

	return strings.Join([]string{"ExternalAccesses", string(data)}, " ")
}
