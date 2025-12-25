package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModuleUpdateRequestPojo 更新模块请求体
type ModuleUpdateRequestPojo struct {

	// 订阅包id
	CloudPackId *string `json:"cloud_pack_id,omitempty"`

	// 订阅包名称
	CloudPackName *string `json:"cloud_pack_name,omitempty"`

	// 订阅包版本
	CloudPackVersion *string `json:"cloud_pack_version,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 布局模块相关信息
	ModuleJson *string `json:"module_json,omitempty"`

	// 名称
	Name string `json:"name"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 模块缩略图
	Thumbnail *string `json:"thumbnail,omitempty"`

	// 模块类型,tab/section
	ModuleType *string `json:"module_type,omitempty"`

	// section类模块关联的指标id
	MetricIds *string `json:"metric_ids,omitempty"`

	// 数据查询方式
	DataQuery *string `json:"data_query,omitempty"`

	// BOA底座版本
	BoaVersion *string `json:"boa_version,omitempty"`
}

func (o ModuleUpdateRequestPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleUpdateRequestPojo struct{}"
	}

	return strings.Join([]string{"ModuleUpdateRequestPojo", string(data)}, " ")
}
