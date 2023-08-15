package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MicroserviceGroup 导入微服务的API分组信息
type MicroserviceGroup struct {

	// 指定已有的分组，为空时创建新的分组
	GroupId *string `json:"group_id,omitempty"`

	// API分组的名称,group_id为空时必填。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头，3-255个字符。 > 中文字符必须为UTF-8或者unicode编码。
	GroupName *string `json:"group_name,omitempty"`

	// group_id为空时必填，指定新分组所属的集成应用
	AppId *string `json:"app_id,omitempty"`
}

func (o MicroserviceGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroserviceGroup struct{}"
	}

	return strings.Join([]string{"MicroserviceGroup", string(data)}, " ")
}
