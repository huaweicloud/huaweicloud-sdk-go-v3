package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpIpGroup ip地址组
type HttpIpGroup struct {

	// ip地址组id
	Id *string `json:"id,omitempty"`

	// ip地址组名称
	Name *string `json:"name,omitempty"`

	// ip地址/地址段列表
	Ips *[]string `json:"ips,omitempty"`

	// ip地址/地址段大小
	Size *int32 `json:"size,omitempty"`

	// ip地址组描述
	Description *string `json:"description,omitempty"`

	// ip地址组创建时间戳
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o HttpIpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpIpGroup struct{}"
	}

	return strings.Join([]string{"HttpIpGroup", string(data)}, " ")
}
