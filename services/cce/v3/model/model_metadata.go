package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 插件基本信息，集合类的元素类型，包含一组由不同名称定义的属性。
type Metadata struct {

	// 唯一id标识
	Uid *string `json:"uid,omitempty" xml:"uid"`

	// 插件名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 插件标签，key/value对格式，接口保留字段，填写不会生效
	Labels map[string]string `json:"labels,omitempty" xml:"labels"`

	// 插件注解，由key/value组成 - 安装：固定值为{\"addon.install/type\":\"install\"} - 升级：固定值为{\"addon.upgrade/type\":\"upgrade\"}
	Annotations map[string]string `json:"annotations,omitempty" xml:"annotations"`

	// 更新时间
	UpdateTimestamp *string `json:"updateTimestamp,omitempty" xml:"updateTimestamp"`

	// 创建时间
	CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
