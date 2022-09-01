package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceUpdateParam struct {

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 实例名。 可以输入中文、数字、字母、下划线、点、破折号。长度介于3-100之间
	DisplayName string `json:"display_name" xml:"display_name"`

	// 自动休眠时长。 arm架构,自动休眠时长只能设置成30，60。x86架构可取值为30，60，240，1440和-1。除-1外，其它值的单位为“分钟”。实例无操作超过自动休眠时长后，将会被暂停（已保存的数据不会被删除）。-1表示实例不会自动停止
	RefreshInterval string `json:"refresh_interval" xml:"refresh_interval"`
}

func (o InstanceUpdateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceUpdateParam struct{}"
	}

	return strings.Join([]string{"InstanceUpdateParam", string(data)}, " ")
}
