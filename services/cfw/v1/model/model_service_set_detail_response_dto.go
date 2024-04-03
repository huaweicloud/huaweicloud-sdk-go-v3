package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceSetDetailResponseDto Query Service Set Detail Response Body
type ServiceSetDetailResponseDto struct {

	// 服务组id
	Id *string `json:"id,omitempty"`

	// 服务组名称
	Name string `json:"name"`

	// 服务组描述信息
	Description *string `json:"description,omitempty"`

	// 服务组类型，0表示自定义服务组，1表示常用WEB服务，2表示常用远程登录和PING，3表示常用数据库
	ServiceSetType *int32 `json:"service_set_type,omitempty"`
}

func (o ServiceSetDetailResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSetDetailResponseDto struct{}"
	}

	return strings.Join([]string{"ServiceSetDetailResponseDto", string(data)}, " ")
}
