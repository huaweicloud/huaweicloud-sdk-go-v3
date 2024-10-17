package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EcsSpecificationBean struct {

	// ECS规格所在的可用区集合
	Azs []string `json:"azs"`

	// 规格ID
	Id string `json:"id"`

	// 规格等级，支持的等级以局点配置为准。 - entry:入门版 - low:基础版 - medium:专业版 - high:高级版
	Level string `json:"level"`

	// 规格名称
	Name string `json:"name"`

	// 规格可添加的数据库数量
	Proxy int32 `json:"proxy"`

	// 内存
	Ram int32 `json:"ram"`

	// CPU
	Vcpus int32 `json:"vcpus"`

	// 可用区类型 - DEDICATED - DEC - EDGE
	AzType *string `json:"az_type,omitempty"`
}

func (o EcsSpecificationBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EcsSpecificationBean struct{}"
	}

	return strings.Join([]string{"EcsSpecificationBean", string(data)}, " ")
}
