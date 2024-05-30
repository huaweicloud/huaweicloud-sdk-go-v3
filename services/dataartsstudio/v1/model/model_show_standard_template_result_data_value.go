package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStandardTemplateResultDataValue 查询数据标准模板结果
type ShowStandardTemplateResultDataValue struct {

	// 数据标准全部属性，集合中是单个StandElementFieldVO对象
	AllFields *[]StandElementFieldVo `json:"allFields,omitempty"`

	// 可选项,集合中是单个StandElementFieldVO对象
	Optional *[]StandElementFieldVo `json:"optional,omitempty"`

	// 系统默认项，集合中是单个StandElementFieldVO对象
	SystemDefault *[]StandElementFieldVo `json:"system_default,omitempty"`

	// 自定义项,集合中是单个StandElementFieldVO对象
	Custom *[]StandElementFieldVo `json:"custom,omitempty"`

	// 是否使用模板
	HasTemplate *bool `json:"hasTemplate,omitempty"`
}

func (o ShowStandardTemplateResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStandardTemplateResultDataValue struct{}"
	}

	return strings.Join([]string{"ShowStandardTemplateResultDataValue", string(data)}, " ")
}
