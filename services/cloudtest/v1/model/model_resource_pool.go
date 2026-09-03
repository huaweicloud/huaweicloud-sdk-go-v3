package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcePool struct {

	// 资源池Id，商用版本使用，数据库中对应label字段
	Id *string `json:"id,omitempty"`

	// 资源池类型，商用版本使用，数据库中对应labelName字段
	Name *string `json:"name,omitempty"`

	// 资源池类型，商用版本使用，数据库中对应labelType字段
	Type *string `json:"type,omitempty"`
}

func (o ResourcePool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcePool struct{}"
	}

	return strings.Join([]string{"ResourcePool", string(data)}, " ")
}
