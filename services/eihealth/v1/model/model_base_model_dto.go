package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaseModelDto struct {

	// 模型名称
	Name *string `json:"name,omitempty"`

	// 模型ID
	Id *string `json:"id,omitempty"`

	// 模型创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 模型描述信息
	Description *string `json:"description,omitempty"`
}

func (o BaseModelDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseModelDto struct{}"
	}

	return strings.Join([]string{"BaseModelDto", string(data)}, " ")
}
